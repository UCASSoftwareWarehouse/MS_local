package project

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mongodb/action/binary"
	"MS_Local/mongodb/action/code"
	"MS_Local/mongodb/model"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	model2 "MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	mongodb2 "MS_Local/utils/mongodb"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Upload(stream pb_gen.MSLocal_UploadServer) error {
	return uploader.doUpload(stream)
}

var uploader = &Uploader{}

type Uploader struct {
}

func (u *Uploader) doUpload(stream pb_gen.MSLocal_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive project info")
	}
	metadata := req.GetMetadata()
	fpath, err := u.receiveStream(stream, metadata)
	if err != nil {
		log.Printf("upload chunk, failed receive stream")
		return err
	}

	if metadata.GetFileType() == pb_gen.FileType_binary {
		log.Printf("receive %d's project %s", metadata.GetProjectId(), metadata.ProjectName)
		u.SaveBinary(fpath, metadata)
	} else if metadata.GetFileType() == pb_gen.FileType_codes {
		log.Printf("receive %d's project %s", metadata.GetProjectId(), metadata.ProjectName)
		u.SaveCodes(fpath, metadata)
	}
	os.Remove(fpath)
	return nil
}

func (u *Uploader) receiveStream(stream pb_gen.MSLocal_UploadServer, metadata *pb_gen.UploadMetadata) (string, error) {
	fileInfo := metadata.GetFileInfo()
	fo, err := os.CreateTemp(config.TempFilePath, fmt.Sprintf("temp_%s", fileInfo.GetFileName()))
	if err != nil {
		log.Printf("create temp file fail, err=[%v]", err)
		stream.SendAndClose(&pb_gen.UploadResponse{
			ProjectInfo: nil,
			Status:      pb_gen.ResponseStatus_fail,
			Message:     "upload failed",
		})
		return "", err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()
	dataSize := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reseive done, file size is %dkb", dataSize/1024)
			break
		}
		if err != nil {
			log.Printf("receive chunk failed, error=[%v]", err)
			stream.SendAndClose(&pb_gen.UploadResponse{
				ProjectInfo: nil,
				Status:      pb_gen.ResponseStatus_fail,
				Message:     "upload failed",
			})
			return "", err
		}
		chunk := req.GetContent()
		dataSize += len(chunk)
		_, err = fo.Write(chunk)
		if err != nil {
			log.Printf("cannot write chunk data: err=[%v]", err)
			stream.SendAndClose(&pb_gen.UploadResponse{
				ProjectInfo: nil,
				Status:      pb_gen.ResponseStatus_fail,
				Message:     "upload failed",
			})
			return "", err
		}
	}
	fpath := filepath.Join(config.TempFilePath, fo.Name())
	return fpath, nil
}

//文件名必须，创建时间和大小似乎无所谓,不如保存上传时间
func (u *Uploader) SaveBinary(fpath string, metadata *pb_gen.UploadMetadata) error {
	//add to mongodb
	binary_id, err := u.SaveFile(fpath, metadata, pb_gen.FileType_binary)
	if err != nil {
		log.Printf("add binary to database failed: %v", err)
		return err
	}
	//update project
	_, err = project.UpdateProject(mysql.Mysql, metadata.GetProjectId(), map[string]interface{}{model2.ProjectColumns.BinaryAddr: mongodb2.ObjectId2String(*binary_id)})
	log.Print("update project binary address success!")
	if err != nil {
		//log.Printf("add binary to database failed: %v", err)
		return err
	}
	return nil
}

func (u *Uploader) SaveCodes(fpath string, metadata *pb_gen.UploadMetadata) error {
	//fname temp文件
	//temp := fmt.Sprintf("extracted_%d", time.Now().UnixNano())
	//tempDir, err := os.MkdirTemp(config.TempFilePath, fmt.Sprintf("extracted_%d", time.Now().UnixNano()))
	//tempDir, err := os.MkdirTemp(config.TempFilePath, temp)
	tempDir, err := os.MkdirTemp(config.TempFilePath, "extracted_")
	if err != nil {
		log.Printf("create temp dir error, err=%v", err)
		return err
	}
	defer os.RemoveAll(tempDir)
	log.Printf("savecodes:create temp dir success!")
	// unzip
	codespath, err := utils.Unzip(fpath, tempDir)
	if err != nil {
		log.Printf("unzip failed, err=[%v]", err)
		return err
	}

	//upload to mongodb code
	code_id, err := u.SaveDir(codespath, metadata)
	if err != nil {
		log.Printf("upload zip to database failed, err=[%v]", err)
		return err
	}
	//update project
	_, err = project.UpdateProject(mysql.Mysql, metadata.GetProjectId(), map[string]interface{}{model2.ProjectColumns.CodeAddr: mongodb2.ObjectId2String(*code_id)})
	log.Print("update project codes address success!")
	if err != nil {
		log.Printf("add binary to database failed: %v", err)
		return err
	}
	return nil
}

func (u *Uploader) SaveDir(dirpath string, metadata *pb_gen.UploadMetadata) (*primitive.ObjectID, error) {
	files, err := os.ReadDir(dirpath)
	dinfo, _ := os.Stat(dirpath)
	if err != nil {
		log.Printf("read dir failed, err=%v", err)
	}
	childFiles := []primitive.ObjectID{}
	var cid *primitive.ObjectID
	for _, file := range files {
		tmpPath := filepath.Join(dirpath, file.Name())
		if file.IsDir() {
			cid, err = u.SaveDir(tmpPath, metadata)
		} else {
			cid, err = u.SaveFile(filepath.Join(dirpath, file.Name()), metadata, pb_gen.FileType_code)
		}
		if err != nil {
			log.Printf("save code file to db error, err=[%v]", err)
		}
		childFiles = append(childFiles, *cid)
	}

	temp_pid, err := code.AddCode(mongodb.CodeCol, model.Code{
		FileName:   dinfo.Name(),
		ProjectID:  metadata.ProjectId,
		FileType:   0, //dir
		FileSize:   uint64(dinfo.Size()),
		Content:    nil,
		UpdateTime: mongodb2.Time2Timestamp(dinfo.ModTime()),
		ChildFiles: childFiles,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("save dir(%s) to mongodb success!", dinfo.Name())
	return &temp_pid, nil
}

// save binary/code file to mongodb
func (u *Uploader) SaveFile(fpath string, metadata *pb_gen.UploadMetadata, filetype pb_gen.FileType) (*primitive.ObjectID, error) {
	content, err := os.ReadFile(fpath)
	if err != nil {
		log.Printf("read temp file failed, %v", err)
		return nil, err
	}
	finfo, _ := os.Stat(fpath)
	var temp_id primitive.ObjectID
	if filetype == pb_gen.FileType_binary {
		//binary　时间来自于包而不是来自于临时文件
		binfo := metadata.GetFileInfo() // binary file info come from request
		temp_id, err = binary.AddBinary(mongodb.BinaryCol, model.Binary{
			FileName:  binfo.FileName,
			ProjectID: metadata.ProjectId,
			FileSize:  binfo.FileSize,
			UpdateTime: primitive.Timestamp{
				T: uint32(binfo.Updatetime.Seconds),
				I: uint32(binfo.Updatetime.Nanos),
			},
			Content: content,
		})
	} else if filetype == pb_gen.FileType_code {
		//code file info come from file itself
		temp_id, err = code.AddCode(mongodb.CodeCol, model.Code{
			FileName:   finfo.Name(),
			ProjectID:  metadata.ProjectId,
			FileSize:   uint64(finfo.Size()),
			FileType:   1, //code file
			UpdateTime: mongodb2.Time2Timestamp(finfo.ModTime()),
			Content:    content,
		})
	}

	if err != nil {
		//log.Printf("add file to database failed: %v, filetype: %v", err, filetype)
		return nil, err
	}
	log.Printf("add file(%s) to mongodb success", finfo.Name())

	return &temp_id, nil
}
