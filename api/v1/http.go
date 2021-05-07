package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

// http 性能相关接口
func GetHttpList(context *gin.Context) {
	var httpParams request.HttpParams
	err := context.BindQuery(&httpParams)
	httpList, err := services.GetHttpInfoList(httpParams.MonitorId, httpParams.StartTime, httpParams.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(httpList, "获取成功", context)
	}
}

func GetHttpStage(context *gin.Context) {
	var queryPageHttp request.HttpParams
	_ = context.BindQuery(&queryPageHttp)
	stageTime, err := services.GetHttpStage(queryPageHttp.MonitorId, queryPageHttp.TimeGrain, queryPageHttp.StartTime, queryPageHttp.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(stageTime, "获取成功", context)
	}
}

func GetHttpQuota(context *gin.Context) {
	var queryPageHttp request.HttpParams
	_ = context.BindQuery(&queryPageHttp)
	quota, err := services.GetHttpQuota(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(quota, "获取成功", context)
	}
}

//
//func GetHttpInfo(context *gin.Context) {
//	var queryPageHttp request.QueryPageHttp
//	err := context.BindQuery(&queryPageHttp)
//	HttpInfoListResponse, err := services.GetHttpInfoList(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
//	HttpQuotaResponse, err := services.GetHttpQuota(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
//	if err != nil {
//		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
//	} else {
//		response.OkWithDetailed(response.PageHttpResponse{
//			HttpListResponse:  HttpInfoListResponse,
//			HttpQuotaResponse: HttpQuotaResponse,
//		}, "获取成功", context)
//	}
//}
//
//func GetHttpStage(context *gin.Context) {
//	var queryPageHttp request.QueryPageHttp
//	_ = context.BindQuery(&queryPageHttp)
//	HttpStageTimeResponse, err := services.GetHttpStageTimeByTimeGrain(queryPageHttp.MonitorId, queryPageHttp.TimeGrain, queryPageHttp.StartTime, queryPageHttp.EndTime)
//	if err != nil {
//		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
//	} else {
//		response.OkWithDetailed(response.PageHttpStage{
//			HttpStageTimeResponse: HttpStageTimeResponse,
//		}, "获取成功", context)
//	}
//}
//
//func GetHttpErrorInfo(context *gin.Context) {
//	var queryPageHttp request.QueryPageHttp
//	_ = context.BindQuery(&queryPageHttp)
//	httpErrorResponse, err := services.GetHttpErrorInfo(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
//	if err != nil {
//		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
//	} else {
//		response.OkWithDetailed(httpErrorResponse, "获取成功", context)
//	}
//}
