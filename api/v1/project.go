package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"dancin-api/utils"
	"github.com/gin-gonic/gin"
)

func GetProjectList(context *gin.Context) {
	emptyList := make([]int, 0)
	claims, exists := context.Get("claims")
	if exists {
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		projectList, err := services.GetProjectList(customClaims.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), context)
			return
		}
		if len(projectList) == 0 {
			response.OkWithDetailed(emptyList, "获取成功", context)
			return
		} else {
			response.OkWithDetailed(projectList, "获取成功", context)
			return
		}
	}
}

func GetProject(context *gin.Context) {
	var projectParams request.ProjectParams
	err := context.ShouldBind(&projectParams)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	project, err := services.GetProject(projectParams.MonitorId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(project, "获取成功", context)
}

func DelProject(context *gin.Context) {
	id, isExist := context.GetQuery("id")
	if isExist {
		err := services.DelProject(id)
		if err != nil {
			response.FailWithMessage(err.Error(), context)
			return
		} else {
			response.OkWithMessage("删除成功！", context)
			return
		}
	} else {
		response.FailWithMessage("不存在项目, 删除失败", context)
		return
	}
}

func GetHealthStatus(context *gin.Context) {
	var healthyParams request.HealthyParams
	emptyList := make([]int, 0)
	err := context.ShouldBind(&healthyParams)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	startTime, endTime := utils.GetTodayStartAndEndTime()
	claims, exists := context.Get("claims")
	if exists {
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		healthStatistics, err := services.GetHealthStatus(customClaims.ID, startTime, endTime)
		if err != nil {
			response.FailWithMessage(err.Error(), context)
			return
		}
		if len(healthStatistics) == 0 {
			response.OkWithDetailed(emptyList, "获取成功", context)
			return
		} else {
			response.OkWithDetailed(healthStatistics, "获取成功", context)
			return
		}
	}
}
