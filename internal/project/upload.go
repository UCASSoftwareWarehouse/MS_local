package project

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mongodb/action"
	"MS_Local/mongodb/action/binary"
	"MS_Local/mongodb/action/code"
	"MS_Local/mongodb/model"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	model2 "MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	mongodb2 "MS_Local/utils/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Upload(stream pb_gen.MSLocal_UploadServer) error {
	return uploader.doUpload(stream)
}

var uploader = &Uploader{}

type Uploader struct {
}

func (u *Uploader) doUpload(stream pb_gen.MSLocal_UploadServer) error {
	log.Println("UPLOAD: start")
	res := &pb_gen.UploadResponse{
		ProjectInfo: nil,
	}
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	metadata := req.GetMetadata()
	fpath, err := u.receiveStream(stream, metadata)
	if err != nil {
		return err
	}

	if metadata.FileInfo.FileType == pb_gen.FileType_binary {
		log.Printf("UPLOAD: receive %d's project %d, upload binary", metadata.UserId, metadata.ProjectId)
		err = u.SaveBinary(fpath, metadata)
	} else if metadata.FileInfo.FileType == pb_gen.FileType_codes {
		log.Printf("UPLOAD: receive %d's project %d, upload codes", metadata.UserId, metadata.ProjectId)
		err = u.SaveCodes(fpath, metadata)
	}
	if err != nil {
		return err
	}
	if fpath != "" {
		os.Remove(fpath)
	}
	pro, err := project.GetProjectById(mysql.Mysql, metadata.ProjectId)
	if err != nil {
		return err
	}
	res.ProjectInfo = &pb_gen.Project{
		Id:          pro.ID,
		ProjectName: pro.ProjectName,
		UserId:      pro.UserID,
		Tags:        pro.Tags,
		License:     pro.License,
		Updatetime: &timestamppb.Timestamp{
			Seconds: pro.UpdateTime.Unix(),
			Nanos:   0,
		},
		ProjectDescription: pro.ProjectDescription,
		CodeAddr:           pro.CodeAddr,
		BinaryAddr:         pro.BinaryAddr,
	}
	err = stream.SendAndClose(res)
	if err != nil {
		log.Printf("cannot send response: err=[%v]", err)
		return err
	}
	log.Printf("UPLOAD: finished.")
	return nil
}

func (u *Uploader) receiveStream(stream pb_gen.MSLocal_UploadServer, metadata *pb_gen.UploadMetadata) (string, error) {
	//fileInfo := metadata.GetFileInfo()
	//log.Printf("file info %v", fileInfo)
	//fo, err := os.CreateTemp(config.TempFilePath, fmt.Sprintf("temp_%s", fileInfo.GetFileName()))
	fo, err := os.CreateTemp(config.Conf.TempFilePath, "temp_upload_")
	if err != nil {
		log.Printf("create temp file fail, err=[%v]", err)
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
			return "", err
		}
		chunk := req.GetContent()
		dataSize += len(chunk)
		_, err = fo.Write(chunk)
		if err != nil {
			log.Printf("cannot write chunk data: err=[%v]", err)
			return "", err
		}
	}
	//fpath := filepath.Join(config.TempFilePath, fo.Name())
	//return fpath, nil
	return fo.Name(), nil //name is path
}

//文件名必须，创建时间和大小似乎无所谓,不如保存上传时间
func (u *Uploader) SaveBinary(fpath string, metadata *pb_gen.UploadMetadata) error {
	//delete old
	DeleteBinary(metadata.ProjectId)
	//add to mongodb
	binary_id, err := u.SaveFile(fpath, metadata, pb_gen.FileType_binary)
	if err != nil {
		log.Printf("add binary to database failed: %v", err)
		return err
	}

	//update project
	err = project.UpdateProject(mysql.Mysql, metadata.GetProjectId(), map[string]interface{}{model2.ProjectColumns.BinaryAddr: mongodb2.ObjectId2String(*binary_id)})
	if err != nil {
		return err
	}
	log.Print("update project binary address success!")
	return nil
}

func (u *Uploader) SaveCodes(fpath string, metadata *pb_gen.UploadMetadata) error {
	//fname temp文件
	//temp := fmt.Sprintf("extracted_%d", time.Now().UnixNano())
	//tempDir, err := os.MkdirTemp(config.TempFilePath, fmt.Sprintf("extracted_%d", time.Now().UnixNano()))
	//tempDir, err := os.MkdirTemp(config.TempFilePath, temp)

	//delete old
	DeleteCodes(metadata.ProjectId)

	tempDir, err := os.MkdirTemp(config.Conf.TempFilePath, "extracted_upload")
	if err != nil {
		log.Printf("create temp dir error, err=%v", err)
		return err
	}
	defer func() {
		err = os.RemoveAll(tempDir)
		if err!=nil{
			log.Printf("delete temp dir  error, err=[%v]", err)
		}
	}()
	log.Printf("savecodes:create temp dir success!")
	//create dir for zip file
	zipname := metadata.FileInfo.FileName
	if strings.HasSuffix(zipname, ".zip"){
		zipname = strings.Replace(zipname, ".zip", "", len(zipname)-4)
	}
	zpath :=filepath.Join(tempDir, zipname)
	err = os.Mkdir(zpath, os.ModePerm)
	if err != nil {
		log.Printf("create zip file path error, err=[%v]", err)
		return err
	}
	// unzip
	codespath, err := utils.Unzip(fpath, zpath)
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
	err = project.UpdateProject(mysql.Mysql, metadata.GetProjectId(), map[string]interface{}{model2.ProjectColumns.CodeAddr: mongodb2.ObjectId2String(*code_id)})
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
			cid, err = u.SaveFile(filepath.Join(dirpath, file.Name()), metadata, pb_gen.FileType_code_file)
		}
		if err != nil {
			log.Printf("save code file to db error, err=[%v]", err)
		}
		childFiles = append(childFiles, *cid)
	}
	//pb_gen.FileType_code_dir
	temp_pid, err := code.AddCode(context.Background(), mongodb.CodeCol, &model.Code{
		FileName:   filepath.Base(dinfo.Name()),
		ProjectID:  metadata.ProjectId,
		FileType:   0, //dir
		FileSize:   uint64(dinfo.Size()),
		ContentID:    "",
		UpdateTime: mongodb2.Time2Timestamp(dinfo.ModTime()),
		ChildFiles: childFiles,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("save dir(%s) to mongodb success!", dinfo.Name())
	return temp_pid, nil
}

// save binary/code file to mongodb
func (u *Uploader) SaveFile(fpath string, metadata *pb_gen.UploadMetadata, filetype pb_gen.FileType) (*primitive.ObjectID, error) {
	//content, err := os.ReadFile(fpath)
	//if err != nil {
	//	log.Printf("read temp file failed, %v", err)
	//	return nil, err
	//}


	var temp_id *primitive.ObjectID
	if filetype == pb_gen.FileType_binary {
		bfile := new(model.Binary)
		binfo := metadata.GetFileInfo() // binary file info come from request
		contentId, err := action.UploadGridFile(fpath, binfo.FileName)
		if(err!=nil){
			return nil, err
		}
		bfile.ContentID = mongodb2.ObjectId2String(*contentId)
		bfile.FileName = binfo.FileName
		bfile.ProjectID = metadata.ProjectId
		//bfile.Content = content
		temp_id, err = binary.AddBinary(context.Background(), mongodb.BinaryCol, bfile)
		if err != nil {
			return nil, err
		}

	} else if filetype == pb_gen.FileType_code_file {
		//code file info come from file itself
		finfo, _ := os.Stat(fpath)
		contentId, err := action.UploadGridFile(fpath, finfo.Name())
		if(err!=nil){
			return nil, err
		}
		temp_id, err = code.AddCode(context.Background(), mongodb.CodeCol, &model.Code{
			FileName:   filepath.Base(finfo.Name()),
			ProjectID:  metadata.ProjectId,
			FileSize:   uint64(finfo.Size()),
			FileType:   1, //code file
			UpdateTime: mongodb2.Time2Timestamp(finfo.ModTime()),
			ContentID: mongodb2.ObjectId2String(*contentId),
		})
		if err != nil {
			return nil, err
		}
	}
	return temp_id, nil
}
