package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHttpInfo(context *gin.Context) {
	startTime, endTime := getTodayStartAndEndTime()
	HttpInfoListResponse, err := services.GetHttpInfoList(startTime, endTime)
	HttpQuotaResponse, err := services.GetHttpQuota(startTime, endTime)
	if err != nil {
		//global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.WebHttpInfoResponse{
			HttpInfoListResponse: HttpInfoListResponse,
			HttpQuotaResponse: HttpQuotaResponse,
		}, "获取成功", context)
	}
}


