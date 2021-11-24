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
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
)

func Download(req *pb_gen.DownloadRequest, stream pb_gen.MSLocal_DownloadServer) error {
	var fpath string
	defer os.Remove(fpath)
	var err error
	if req.FileType == pb_gen.FileType_binary {
		fpath, err = DownloadBinary(req.FileId, "")
	} else if req.FileType == pb_gen.FileType_code {
		fpath, err = DownloadCode(req.FileId, "")
	} else if req.FileType == pb_gen.FileType_project {
		fpath, err = DownloadProject(req.FileId)
	} else {
		log.Printf("download unrecognized file type")
		return status.Errorf(codes.InvalidArgument, "download unrecognized file type")
	}
	if err != nil {
		return err
	}
	//1 判断fpath是文件名还是路径名
	//2 send
	return nil
}

func DownloadBinary(fid string, fpath string) (string, error) {
	//search file
	//binaryfile, err := binary.GetBinaryByFileId(stream.Context(), mongodb.BinaryCol, mongodb2.String2ObjectId(fid))
	binaryfile, err := binary.GetBinaryByFileId(context.Background(), mongodb.BinaryCol, mongodb2.String2ObjectId(fid))
	if err != nil {
		return "", err
	}
	var fo *os.File

	if fpath == "" {
		fo, err = os.CreateTemp(config.TempFilePath, "temp_dbin_")
	} else {
		fo, err = os.Create(filepath.Join(fpath, binaryfile.FileName))
	}

	if err != nil {
		log.Printf("create temp file fail, err=[%v]", err)
		return "", err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()

	_, err = fo.Write(binaryfile.Content)
	if err != nil {
		log.Printf("wirte file error")
		return "", err
	}
	log.Printf("wirte binary to %v", fo.Name())
	return fo.Name(), nil
}

func DownloadCode(fid string, fpath string) (string, error) {
	//codefile, err := code2.GetCodeByFileId(stream.Context(), mongodb.CodeCol, mongodb2.String2ObjectId(fid))
	codefile, err := code2.GetCodeByFileId(context.Background(), mongodb.CodeCol, mongodb2.String2ObjectId(fid))
	if err != nil {
		return "", err
	}
	if codefile.FileType == 0 { //dir
		cdir := filepath.Join(fpath, codefile.FileName)
		err = os.Mkdir(cdir, os.ModePerm)
		if err != nil {
			log.Printf("mkdir %s failed", cdir)
			return "", nil
		}
		for _, cid := range codefile.ChildFiles {
			_, err = DownloadCode(mongodb2.ObjectId2String(cid), cdir)
			if err != nil {
				log.Printf("download code failed, id = [%v]", cid)
				return "", err
			}
		}
		log.Printf("download code dir(%s) usccess", codefile.FileName)
		return cdir, nil
	} else if codefile.FileType == 1 { //file
		var fo *os.File
		if fpath == "" {
			fo, err = os.CreateTemp(config.TempFilePath, "temp_dcode_")
		} else {
			fo, err = os.Create(filepath.Join(fpath, codefile.FileName))
		}

		if err != nil {
			log.Printf("create temp file fail, err=[%v]", err)
			return "", err
		}
		defer func() {
			if err := fo.Close(); err != nil {
				log.Printf("Upload close fo failed, err=[%+v]", err)
			}
		}()

		_, err = fo.Write(codefile.Content)
		if err != nil {
			log.Printf("wirte file error")
			return "", err
		}
		log.Printf("wirte code to %v", fo.Name())
		return fo.Name(), err
	}
	return "", status.Errorf(codes.InvalidArgument, "no such filetype")
}

func DownloadProject(pid string) (string, error) {
	tempDir, err := os.MkdirTemp(config.TempFilePath, "download_tempdir_")
	if err != nil {
		log.Printf("create temp dir error, err=%v", err)
		return "", err
	}
	defer os.RemoveAll(tempDir)
	pid_ := utils.String2Uint64(pid)
	project, err := project.GetProjectById(mysql.Mysql, pid_)
	if err != nil {
		return "", err
	}
	//download binary to temp
	_, err = DownloadBinary(project.BinaryAddr, tempDir)
	if err != nil {
		return "", err
	}
	//download codes
	_, err = DownloadCode(project.CodeAddr, tempDir)
	if err != nil {
		return "", err
	}

	fpath, err := utils.Zip(tempDir, "")
	if err != nil {
		log.Printf("zip failed, err=[%v]", err)
	}
	return fpath, nil
}
