package services

import (
	"danci-api/model/response"
)

func GetProjectList(id uint) (projectList []response.ProjectResponse) {
	//global.GVA_DB.Model(&model.TeamProject{}).Select("project").Where("member = ?", id).Scan(&teamProjectIds)
	//fmt.Println(teamProjectIds, "teamProjectIds")
	//global.GVA_DB.Model(&model.Project{}).Where("id in ?", teamProjectIds).Find(&projectList)
	return
}
