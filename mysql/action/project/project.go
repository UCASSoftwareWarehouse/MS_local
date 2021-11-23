package project

import (
	"MS_Local/mysql/model"
	"gorm.io/gorm"
	"log"
)

// Add Project
func AddProject(db *gorm.DB, project model.Project) (uint64, error) {
	err := db.Create(&project).Error
	if err != nil {
		log.Printf("add project error, err=[%v]", err)
	}
	return project.ID, err
}

func UpdateProject(db *gorm.DB, projectId uint64, options map[string]interface{}) (model.Project, error) {
	project, _ := GetProjectById(db, projectId)
	err := db.Model(&project).Updates(options).Error
	if err != nil {
		log.Printf("Update project failed, err=%v", err)
	}
	return project, err
}

// get all software project
func GetAllMetadata(db *gorm.DB) ([]model.Project, error) {
	var projects []model.Project
	result := db.Find(&projects)
	if result.Error != nil {
		log.Printf("get all project error, err=[%v]", result.Error)
	}
	return projects, result.Error
}

// get softmatadata by primary key projectid
func GetProjectById(db *gorm.DB, projectId uint64) (model.Project, error) {
	project := new(model.Project)
	err := db.Where("ID = ?", projectId).First(project).Error
	if err != nil {
		log.Printf("get project by project id error, err=[%v]", err)
	}
	return *project, err
}

// Get all software project by user id
func GetProjectsByUserId(db *gorm.DB, userId uint64) ([]model.Project, error) {
	var projects []model.Project
	err := db.Where("user_id = ?", userId).Find(&projects).Error
	if err != nil {
		log.Printf("get project by user id error, err=[%v]", err)
	}
	return projects, err
}

// get by projectName, userid
func GetProjectsByOptions(db *gorm.DB, options map[string]interface{}) ([]model.Project, error) {
	var projects []model.Project
	err := db.Where(options).Find(&projects).Error
	if err != nil {
		log.Printf("get project by options error, err=[%v]", err)
	}
	return projects, err
}

// Delete Project from list
func DeleteProjects(db *gorm.DB, projects []model.Project) error {
	for _, project := range projects {
		err := db.Delete(&project).Error
		if err != nil {
			log.Printf("delete project error, err=[%v]", err)
			return err
		}
	}
	return nil
}

func DeleteProjectById(db *gorm.DB, id uint64) error {
	project, _ := GetProjectById(db, id)
	err := db.Delete(&project).Error
	if err != nil {
		log.Printf("delete project error, err=[%v]", err)
	}
	return err
}