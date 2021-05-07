package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetResourceErrorInfo(context *gin.Context) {
	var resourceErrorParams request.ResourceErrorParams
	err := context.BindQuery(&resourceErrorParams)
	startTime, endTime := getTodayStartAndEndTime()
	ResourcesList, err := services.GetResourcesInfoList(resourceErrorParams.MonitorId, startTime, endTime)
	ResourcesQuota, err := services.GetResourcesQuota(resourceErrorParams.MonitorId, startTime, endTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.PageResourcesResponse{
			ResourcesList:  ResourcesList,
			ResourcesQuota: ResourcesQuota,
		}, "获取成功", context)
	}
}
