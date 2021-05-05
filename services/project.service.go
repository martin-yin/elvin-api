package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetProjectList(id uint) (projectList []response.ProjectResponse) {
	//global.GVA_DB.Model(&model.TeamProject{}).Select("project").Where("member = ?", id).Scan(&teamProjectIds)
	//fmt.Println(teamProjectIds, "teamProjectIds")
	//global.GVA_DB.Model(&model.Project{}).Where("id in ?", teamProjectIds).Find(&projectList)
	return
}

func CreateProject(project model.Project) (projectInter model.Project, err error) {
	err = global.GVA_DB.Model(&model.Project{}).Create(&project).Error
	return project, err
}
