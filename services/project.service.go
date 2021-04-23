package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func GetProjectList() (projectList []response.ProjectResponse) {
	global.GVA_DB.Model(&model.Project{}).Find(&projectList)
	fmt.Println(projectList, "projectList!!")
	return projectList
}

