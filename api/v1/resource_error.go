package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetResourceErrorInfo(context *gin.Context) {
	startTime, endTime := getTodayStartAndEndTime()
	ResourcesList, err := services.GetResourcesInfoList(startTime, endTime)
	ResourcesQuota, err := services.GetResourcesQuota(startTime, endTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.PageResourcesResponse{
			ResourcesList:  ResourcesList,
			ResourcesQuota: ResourcesQuota,
		}, "获取成功", context)
	}
}
