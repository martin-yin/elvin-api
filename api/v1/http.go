package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHttpInfo(context *gin.Context) {
	var queryPageHttp request.QueryPageHttp
	err := context.BindQuery(&queryPageHttp)
	HttpInfoListResponse, err := services.GetHttpInfoList(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
	HttpQuotaResponse, err := services.GetHttpQuota(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.PageHttpResponse{
			HttpListResponse:  HttpInfoListResponse,
			HttpQuotaResponse: HttpQuotaResponse,
		}, "获取成功", context)
	}
}

func GetHttpStage(context *gin.Context) {
	var queryPageHttp request.QueryPageHttp
	_ = context.BindQuery(&queryPageHttp)

	if queryPageHttp.StageType == "success" {
		HttpStageTimeResponse, err := services.GetHttpStageTimeSuccess(queryPageHttp.StartTime, queryPageHttp.EndTime, queryPageHttp.TimeGrain)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		} else {
			response.OkWithDetailed(response.PageHttpStage{
				HttpStageTimeResponse: HttpStageTimeResponse,
			}, "获取成功", context)
		}
	}
}
