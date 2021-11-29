package project

import (
	"MS_Local/mysql/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func UpdateProject(db *gorm.DB, projectId uint64, options map[string]interface{}) error {
	//project, _ := GetProjectById(db, projectId)
	//err := db.Model(&project).Updates(options).Error
	err := db.Model(&model.Project{}).Where("ID = ?", projectId).Updates(options).Error
	if err != nil {
		log.Printf("Update project failed, err=%v", err)
	}

	return err
}

// get all software project
func GetLimitProjects(db *gorm.DB, num int, page int, projects *[]model.Project) error {
	offset := (page - 1) * num
	//result := db.Offset(offset).Limit(num).Find(&projects)
	filter := db.Offset(offset).Limit(num)
	result := filter.Find(projects)
	if result.Error != nil {
		log.Printf("get all project error, err=[%v]", result.Error)
		return result.Error
	}
	return result.Error
}

// get softmatadata by primary key projectid
func GetProjectById(db *gorm.DB, projectId uint64) (*model.Project, error) {
	project := new(model.Project)
	err := db.Where("ID = ?", projectId).First(project).Error
	if err != nil {
		log.Printf("get project by project id error, err=[%v]", err)
		return nil, err
	}
	log.Printf("get project info by id success!")
	return project, err
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

// Get all software project by user id
func GetProjectsByUserId(db *gorm.DB, userId uint64, limit int, page int, projects *[]model.Project) error {
	offset := (page - 1) * limit
	filter := db.Limit(limit).Where("user_id = ?", userId).Offset(offset)
	err := filter.Order(model.ProjectColumns.ProjectName).Order(model.ProjectColumns.ID + " desc").Find(&projects).Error
	if err != nil {
		log.Printf("get project by user id error, err=[%v]", err)
		return err
	}
	return nil
}

func SearchProjectByName(db *gorm.DB, keyword string, limit int, page int, projects *[]model.Project) error {
	//basedon like
	//pattern := strings.Join([]string{"%", keyword, "%"}, "")
	offset := (page - 1) * limit
	//filter := db.Limit(limit).Where("project_name LIKE ?", pattern).Offset(offset)
	dislimit := 5
	filter := db.Limit(limit).Where("levenshtein(project_name,?)<?", keyword, dislimit).Offset(offset)
	//err := filter.Find(projects).Error
	err := filter.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "levenshtein(project_name,?)", Vars: []interface{}{keyword}, WithoutParentheses: true},
	}).Find(projects).Error
	if err != nil {
		log.Printf("serach failed, err=[%v]", err)
		return err
	}

	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
	//full text 基于词频，不方便使用
	//if err := db.Table("project").Where("MATCH (project_name) AGAINST (? IN BOOLEAN MODE)", keyword).Find(&projects).Error; err != nil {
	//	return err
	//}

	return nil
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
