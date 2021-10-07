package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"dancin-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetResourceError(context *gin.Context) {
	var resourceErrorParams request.ResourceErrorParams
	err := context.BindQuery(&resourceErrorParams)
	startTime, endTime := utils.GetTodayStartAndEndTime()
	ResourcesList, err := services.GetResourcesList(resourceErrorParams.MonitorId.MonitorId, startTime, endTime)
	ResourcesQuota, err := services.GetResourcesQuota(resourceErrorParams.MonitorId.MonitorId, startTime, endTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(response.PageResourcesResponse{
		ResourcesList:  ResourcesList,
		ResourcesQuota: ResourcesQuota,
	}, "获取成功", context)
}
