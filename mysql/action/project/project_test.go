package project

import (
	"MS_Local/config"
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
	config.InitConfig()
	mysql.InitMysql()
	project := model.Project{
		ProjectName:        "projectaaa",
		UserID:             2,
		Tags:               "v2.0.0",
		CodeAddr:           "",
		BinaryAddr:         "",
		License:            "test license",
		ProjectDescription: "test description",
		UpdateTime:         time.Now(),
	}
	id, err := AddProject(mysql.Mysql, project)
	if err != nil {
		t.Errorf("add project error：%s", err)
	}
	fmt.Println(id)
}

func TestUpdateProject(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	temp_id := primitive.NewObjectID()
	string_id := mongodb.ObjectId2String(temp_id)

	err := UpdateProject(mysql.Mysql, 1, map[string]interface{}{model.ProjectColumns.BinaryAddr: string_id})
	if err != nil {
		t.Errorf("update object id error: %v", err)
	}
}

func TestGetProjectsByUserId(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	var projects []model.Project
	err := GetProjectsByUserId(mysql.Mysql, 1, 10, 1, &projects)
	if err != nil {
		t.Errorf("get user's all projects error ：%s", err)
	}
	for _, pro := range projects {
		log.Printf("%v", pro)
	}
}

func TestGetProjectById(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	project, err := GetProjectById(mysql.Mysql, 7)
	if err != nil {
		t.Errorf("get project by options error：%s", err)
	}
	fmt.Println(project)
}
func TestGetProjectsByOptions(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	//test get by project name user id
	project1, err1 := GetProjectsByOptions(mysql.Mysql, map[string]interface{}{"project_name": "test name", "user_id": 1})
	if err1 != nil {
		t.Errorf("get project by project_id & user_id error：%s", err1)
	}
	fmt.Println(project1)
	//test get by project_name, user_id, tag
	project2, err2 := GetProjectsByOptions(mysql.Mysql, map[string]interface{}{"project_name": "test name", "user_id": 1, "tags": "v"})
	if err2 != nil {
		t.Errorf("get project by project_id & user_id & tag error：%s", err2)
	}
	fmt.Println(project2)
}

func TestDeleteProjects(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	var projects []model.Project
	GetProjectsByUserId(mysql.Mysql, 1, 1, 2, &projects)
	err := DeleteProjects(mysql.Mysql, projects)
	if err != nil {
		t.Errorf("delete project failed：%s", err)
	}
}

func TestGetLimitProjects(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	var projects []model.Project
	err := GetLimitProjects(mysql.Mysql, 2, 1, &projects)
	if err != nil {
		t.Errorf("err=[%v]", err)
	}
	log.Println(len(projects))
}

func TestSearchProjectByName(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	var projects []model.Project
	err := SearchProjectByName(mysql.Mysql, "ject", 10, 1, &projects)
	if err != nil {
		t.Errorf("err=[%v]", err)
	}
	for _, pro := range projects {
		log.Printf("%v", pro)
	}
}

func TestDeleteProjectById(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	err := DeleteProjectById(mysql.Mysql, 12)
	if err != nil {
		t.Errorf("delete project by id failed：%s", err)
	}
}

//func TestSearchProjectByName(t *testing.T) {
//	db, _ := mysql.InitMysql()
//	projects, err := SearchProjectByName(db, "")
//	if err != nil {
//		t.Error(err)
//	}
//	log.Println(len(projects))
//}
