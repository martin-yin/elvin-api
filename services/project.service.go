package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func GetProjectList(id uint) (projectList []response.ProjectResponse) {
	var teamProjectIds []string
	//var productIds string
	global.GVA_DB.Model(&model.TeamProject{}).Select("project").Where("member = ?", id).Scan(&teamProjectIds)

	fmt.Println(teamProjectIds, "teamProjectIds")
	//for _, value := range teamProjectIds {
	//	productIds = productIds + value + ","
	//}
	//ids := strings.Split(productIds[0:len(productIds)-1], ",")
	//fmt.Println(ids, "ids")
	global.GVA_DB.Model(&model.Project{}).Where("id in ?", teamProjectIds).Find(&projectList)
	return
}
