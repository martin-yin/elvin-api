package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"fmt"
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
		} else {
			if len(projectList) == 0 {
				response.OkWithDetailed(emptyList, "获取成功", context)
				return
			} else {
				response.OkWithDetailed(projectList, "获取成功", context)
				return
			}
		}
		return
	}
}

func GetProjectHealthy(context *gin.Context) {
	var healthyParams request.HealthyParams
	err := context.ShouldBind(&healthyParams)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	fmt.Print(healthyParams, "healthyParams \n")

	startTime, endTime := getTodayStartAndEndTime()
	healthyData, err := services.GetProjectHealthy(startTime, endTime, healthyParams.MonitorId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	} else {
		response.OkWithDetailed(healthyData, "获取成功", context)
		return
	}
}
