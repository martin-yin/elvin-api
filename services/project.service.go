package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func GetHomeProjectStatisticsData() {
	var projectList []response.HomeStatisticsDataProjectResponse
	fmt.Println(projectList)
}


func GetProjectList() (projectList []response.ProjectResponse){
	global.GVA_DB.Model(&model.Project{}).Find(&projectList)
	return
}