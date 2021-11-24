package project

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mongodb/action/binary"
	code2 "MS_Local/mongodb/action/code"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	mongodb2 "MS_Local/utils/mongodb"
	"bufio"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Download(req *pb_gen.DownloadRequest, stream pb_gen.MSLocal_DownloadServer) error {
	var fpath string
	var err error
	var fminfo *pb_gen.FileInfo
	if req.FileType == pb_gen.FileType_binary {
		fpath, fminfo, err = DownloadBinary(req.FileId, "")
	} else if req.FileType == pb_gen.FileType_code {
		fpath, fminfo, err = DownloadCode(req.FileId, "")
	} else if req.FileType == pb_gen.FileType_project {
		fpath, fminfo, err = DownloadProject(req.FileId)
	} else if req.FileType == pb_gen.FileType_codes {
		fpath, fminfo, err = DownloadCodes(req.FileId, "")
	} else {
		log.Printf("download unrecognized file type")
		return status.Errorf(codes.InvalidArgument, "download unrecognized file type")
	}
	if err != nil {
		return err
	}
	defer func() {
		if err := os.Remove(fpath); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()
	//response metadata
	res := &pb_gen.DownloadResponse{
		Data: &pb_gen.DownloadResponse_Metadata{
			Metadata: &pb_gen.DownloadMetadate{
				FileInfo: fminfo,
				FileType: req.FileType,
				Status:   pb_gen.ResponseStatus_ok,
				Message:  "download success",
			}}}
	err = stream.Send(res)
	if err != nil {
		log.Printf("cannot send metadat info to client:err=[%v]", err)
		return err
	}

	//2 send
	err = SendStream(fpath, stream)
	if err != nil {
		return err
	}
	return nil
}

func SendStream(fpath string, stream pb_gen.MSLocal_DownloadServer) error {
	//send file content
	fo, err := os.Open(fpath)
	if err != nil {
		log.Printf("cannot open file, err=[%v]", err)
		return err
	}
	defer fo.Close()
	reader := bufio.NewReader(fo)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot read chunk to buffer, err=[%v]", err)
			return err
		}
		req := &pb_gen.DownloadResponse{
			Data: &pb_gen.DownloadResponse_Content{
				Content: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Printf("cannot send chunk to server,err=[%v], message=[%v] ", err, stream.RecvMsg(nil))
			return err
		}
	}
	log.Printf("send file done")
	return nil
}

func DownloadBinary(fid string, fpath string) (string, *pb_gen.FileInfo, error) {
	//search file
	//binaryfile, err := binary.GetBinaryByFileId(stream.Context(), mongodb.BinaryCol, mongodb2.String2ObjectId(fid))
	binaryfile, err := binary.GetBinaryByFileId(context.Background(), mongodb.BinaryCol, mongodb2.String2ObjectId(fid))
	if err != nil {
		return "", nil, err
	}
	var fo *os.File

	if fpath == "" {
		fo, err = os.CreateTemp(config.TempFilePath, "temp_dbin_")
	} else {
		fo, err = os.Create(filepath.Join(fpath, binaryfile.FileName))
	}

	if err != nil {
		log.Printf("create temp file fail, err=[%v]", err)
		return "", nil, err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()

	_, err = fo.Write(binaryfile.Content)
	if err != nil {
		log.Printf("wirte file error")
		return "", nil, err
	}
	log.Printf("wirte binary to %v", fo.Name())
	fminfo := &pb_gen.FileInfo{
		FileName: binaryfile.FileName,
		FileSize: binaryfile.FileSize,
		Updatetime: &timestamppb.Timestamp{
			Seconds: int64(binaryfile.UpdateTime.T),
			Nanos:   int32(binaryfile.UpdateTime.I),
		},
	}
	return fo.Name(), fminfo, nil
}

func DownloadCode(fid string, fpath string) (string, *pb_gen.FileInfo, error) {
	codefile, err := code2.GetCodeByFileId(context.Background(), mongodb.CodeCol, mongodb2.String2ObjectId(fid))
	if err != nil {
		return "", nil, err
	}
	fminfo := &pb_gen.FileInfo{
		FileName: codefile.FileName,
		FileSize: codefile.FileSize,
		Updatetime: &timestamppb.Timestamp{
			Seconds: int64(codefile.UpdateTime.T),
			Nanos:   int32(codefile.UpdateTime.I),
		},
	}
	if codefile.FileType == 0 { //dir
		var cdir string
		if fpath == "" {
			cdir, err = os.MkdirTemp(config.TempFilePath, "temp_dcode_")
		} else {
			cdir = filepath.Join(fpath, codefile.FileName)
			err = os.Mkdir(cdir, os.ModePerm)
		}

		if err != nil {
			log.Printf("mkdir %s failed", cdir)
			return "", nil, nil
		}
		for _, cid := range codefile.ChildFiles {
			_, _, err = DownloadCode(mongodb2.ObjectId2String(cid), cdir)
			if err != nil {
				log.Printf("download code failed, id = [%v]", cid)
				return "", nil, err
			}
		}
		log.Printf("download code dir(%s) usccess", codefile.FileName)

		return cdir, fminfo, nil
	} else if codefile.FileType == 1 { //file
		var fo *os.File
		if fpath == "" {
			fo, err = os.CreateTemp(config.TempFilePath, "temp_dcode_")
		} else {
			fo, err = os.Create(filepath.Join(fpath, codefile.FileName))
		}

		if err != nil {
			log.Printf("create temp file fail, err=[%v]", err)
			return "", nil, err
		}
		defer func() {
			if err := fo.Close(); err != nil {
				log.Printf("Upload close fo failed, err=[%+v]", err)
			}
		}()

		_, err = fo.Write(codefile.Content)
		if err != nil {
			log.Printf("wirte file error")
			return "", nil, err
		}
		log.Printf("wirte code to %v", fo.Name())
		return fo.Name(), fminfo, err
	}
	return "", nil, status.Errorf(codes.InvalidArgument, "no such filetype")
}

func DownloadCodes(fid string, fpath string) (string, *pb_gen.FileInfo, error) {
	if fpath == "" {
		dir, err := os.MkdirTemp(config.TempFilePath, "temp_dcodes_")
		if err != nil {
			log.Printf("create temp dir error, err=[%v]", err)
			return "", nil, err
		}
		defer os.RemoveAll(dir)
		codes_dir, fminfo, err := DownloadCode(fid, dir)
		if err != nil {
			return "", nil, err
		}
		//压缩codes
		zpath, err := utils.Zip(codes_dir, "") //temp file
		if err != nil {
			log.Printf("zip failed, err=[%v]", err)
		}
		return zpath, fminfo, nil
	} else { //download到指定文件夹
		codes_dir, fminfo, err := DownloadCode(fid, fpath)
		if err != nil {
			return "", nil, err
		}
		return codes_dir, fminfo, nil
	}
}

func DownloadProject(pid string) (string, *pb_gen.FileInfo, error) {
	tempDir, err := os.MkdirTemp(config.TempFilePath, "download_tempdir_")
	if err != nil {
		log.Printf("create temp dir error, err=%v", err)
		return "", nil, err
	}
	defer os.RemoveAll(tempDir)
	pid_ := utils.String2Uint64(pid)
	project, err := project.GetProjectById(mysql.Mysql, pid_)
	if err != nil {
		return "", nil, err
	}
	//download binary to temp
	_, _, err = DownloadBinary(project.BinaryAddr, tempDir)
	if err != nil {
		return "", nil, err
	}
	//download codes
	_, _, err = DownloadCodes(project.CodeAddr, tempDir)
	if err != nil {
		return "", nil, err
	}

	fpath, err := utils.Zip(tempDir, "")
	if err != nil {
		log.Printf("zip failed, err=[%v]", err)
	}
	fminfo := &pb_gen.FileInfo{
		FileName: project.ProjectName,
		FileSize: 0,
		Updatetime: &timestamppb.Timestamp{
			Seconds: int64(project.UpdateTime.Unix()),
			Nanos:   0,
		},
	}
	return fpath, fminfo, nil
}
