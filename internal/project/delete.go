package project

import (
	"MS_Local/mongodb"
	"MS_Local/mongodb/action"
	"MS_Local/mongodb/action/binary"
	code2 "MS_Local/mongodb/action/code"
	model2 "MS_Local/mongodb/model"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	mongodb2 "MS_Local/utils/mongodb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func Delete(ctx context.Context, req *pb_gen.DeleteProjectRequest) (*pb_gen.DeleteProjectResponse, error) {
	//get

	if req.FileType == pb_gen.FileType_project {
		log.Println("DELETE: delete project")
		err := DeleteProject(req.Pid)
		if err != nil {
			return nil, err
		}
	} else if req.FileType == pb_gen.FileType_binary {
		log.Println("DELETE: delete binary")
		err := DeleteBinary(req.Pid)
		if err != nil {
			return nil, err
		}
	} else if req.FileType == pb_gen.FileType_codes {
		log.Println("DELETE: delete codes")
		err := DeleteCodes(req.Pid)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "can only delete pro/codes/binary")
	}
	//if ctx.Err() == context.Canceled {
	//	log.Print("request is canceled")
	//	return nil, status.Error(codes.Canceled, "request is canceled")
	//}
	//
	//if ctx.Err() == context.DeadlineExceeded {
	//	log.Print("deadline is exceeded")
	//	return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	//}
	log.Printf("DELETE: finish!")
	return &pb_gen.DeleteProjectResponse{
		Message: "success",
	}, nil
}

func DeleteProject(pid uint64) error {
	pro, err := project.GetProjectById(mysql.Mysql, pid)
	if err != nil {
		return err
	}
	//delete binary
	baddr := pro.BinaryAddr
	if baddr != "" {
		err := DeleteBinary(pro.ID)
		if err != nil {
			return err
		}
	}
	//delete codes
	caddr := pro.CodeAddr
	if caddr != "" {
		err := DeleteCodes(pro.ID)
		if err != nil {
			return err
		}
	}
	//delete project
	err = project.DeleteProjectById(mysql.Mysql, pro.ID)
	if err != nil {
		return err
	}
	log.Println("DELETE: delete project success")
	return nil
}

func DeleteBinary(pid uint64) error {
	//check
	pro, err := project.GetProjectById(mysql.Mysql, pid)
	if err != nil {
		return err
	}
	//no binary
	if pro.BinaryAddr == "" {
		return nil
	}
	//delete content
	binfo,err := binary.GetBinaryByFileId(context.Background(), mongodb.BinaryCol, mongodb2.String2ObjectId(pro.BinaryAddr))
	if binfo.ContentID!=""{
		action.DeleteGridFile(binfo.ContentID)
	}

	//delete
	err = binary.DeleteBinaryByProjectId(context.Background(), mongodb.BinaryCol, pid)
	if err != nil {
		return err
	}
	//update
	err = project.UpdateProject(mysql.Mysql, pid, map[string]interface{}{model.ProjectColumns.BinaryAddr: ""})
	if err != nil {
		return err
	}
	log.Println("DELETE: delete %d binary success", pid)
	return nil
}

func DeleteCodes(pid uint64) error {
	//check
	pro, err := project.GetProjectById(mysql.Mysql, pid)
	if err != nil {
		return err
	}
	if pro.CodeAddr == "" {
		return nil
	}

	if pro.CodeAddr==""{
		return nil;
	}
	//delete codes content
	cinfo, err := code2.GetCodeByFileId(context.Background(), mongodb.CodeCol, mongodb2.String2ObjectId(pro.CodeAddr))
	if err!=nil{
		return err;
	}
	var queue []*model2.Code
	queue = append(queue, cinfo)
	for{
		if(len(queue)==0){
			break;
		}
		temp := queue[0]
		queue = queue[1:len(queue)]
		if(temp.FileType==0){//dir
			for _, cid := range(temp.ChildFiles){
				temp_cinfo, err := code2.GetCodeByFileId(context.Background(), mongodb.CodeCol, cid)
				if err!=nil{
					return err
				}
				queue = append(queue, temp_cinfo)
			}
		}else{
			err := action.DeleteGridFile(temp.ContentID)
			if err!=nil{
				return err
			}
		}
	}
	//delete codes metadata
	err = code2.DeleteManyCodesByProjectId(context.Background(), mongodb.CodeCol, pid)
	if err != nil {
		return err
	}
	//update
	err = project.UpdateProject(mysql.Mysql, pid, map[string]interface{}{model.ProjectColumns.CodeAddr: ""})
	if err != nil {
		return err
	}

	log.Println("DELETE: delete %d code success", pid)
	return nil
}
