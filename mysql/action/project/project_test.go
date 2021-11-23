package project

import (
	"MS_Local/mysql"
	"MS_Local/mysql/model"
	"MS_Local/utils/mongodb"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
	"time"
)

func TestAddProject(t *testing.T) {
	db, _ := mysql.InitMysql()
	project := model.Project{
		ProjectName:        "test name",
		UserID:             1,
		Tags:               "v",
		CodeAddr:           "",
		BinaryAddr:         "",
		License:            "test license",
		ProjectDescription: "test description",
		UpdateTime:         time.Now(),
	}
	id, err := AddProject(db, project)
	if err != nil {
		t.Errorf("add project error：%s", err)
	}
	fmt.Println(id)
}
func TestUpdateProject(t *testing.T) {
	db, _ := mysql.InitMysql()
	temp_id := primitive.NewObjectID()
	string_id := mongodb.ObjectId2String(temp_id)

	project, err := UpdateProject(db, 1, map[string]interface{}{model.ProjectColumns.BinaryAddr: string_id})
	if err != nil {
		t.Errorf("update object id error: %v", err)
	}
	log.Println(project)
}

func TestGetProjectsByUserId(t *testing.T) {
	db, _ := mysql.InitMysql()
	projects, err := GetProjectsByUserId(db, 1)
	if err != nil {
		t.Errorf("get user's all projects error ：%s", err)
	}
	fmt.Println(projects)
}

func TestGetProjectById(t *testing.T) {
	db, _ := mysql.InitMysql()
	projects, err := GetProjectById(db, 1)
	if err != nil {
		t.Errorf("get project by options error：%s", err)
	}
	fmt.Println(projects)
}
func TestGetProjectsByOptions(t *testing.T) {
	db, _ := mysql.InitMysql()
	//test get by project name user id
	project1, err1 := GetProjectsByOptions(db, map[string]interface{}{"project_name": "test name", "user_id": 1})
	if err1 != nil {
		t.Errorf("get project by project_id & user_id error：%s", err1)
	}
	fmt.Println(project1)
	//test get by project_name, user_id, tag
	project2, err2 := GetProjectsByOptions(db, map[string]interface{}{"project_name": "test name", "user_id": 1, "tags": "v"})
	if err2 != nil {
		t.Errorf("get project by project_id & user_id & tag error：%s", err2)
	}
	fmt.Println(project2)
}

func TestDeleteProjects(t *testing.T) {
	db, _ := mysql.InitMysql()
	projects, _ := GetProjectsByUserId(db, 1)
	err := DeleteProjects(db, projects)
	if err != nil {
		t.Errorf("delete project failed：%s", err)
	}
}

func TestDeleteProjectById(t *testing.T) {
	db, _ := mysql.InitMysql()
	err := DeleteProjectById(db, 1)
	if err != nil {
		t.Errorf("delete project by id failed：%s", err)
	}
}
