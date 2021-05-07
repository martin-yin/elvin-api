package services

import (
	"danci-api/global"
	"danci-api/model"
)

func GetProjectList(id uint) (projectList []model.Project, err error) {
	var teamList []model.Team
	err = global.GVA_DB.Preload("Admins", "id = ? ", 1).Preload("Projects").Model(&model.Team{}).Find(&teamList).Error
	for _, team := range teamList {
		for _, project := range team.Projects {
			projectList = append(projectList, project)
		}
	}
	return
}


func CreateProject(project model.Project) (projectInter model.Project, err error) {
	err = global.GVA_DB.Model(&model.Project{}).Create(&project).Error
	return project, err
}

func FindProject(projectName string) (isExist bool) {
	var project model.Project
	result := global.GVA_DB.Model(&model.Project{}).Where("project_name = ? ", projectName).First(&project)
	if result.RowsAffected != 0 {
		return true
	}
	return false
}
