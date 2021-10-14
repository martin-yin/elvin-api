package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
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
		return
	}
	response.OkWithDetailed(httpList, "获取成功", context)
}

func GetHttpErrorList(context *gin.Context) {
	var httpParams request.HttpParams
	err := context.BindQuery(&httpParams)
	httpList, err := services.GetHttpErrorList(httpParams.MonitorId, httpParams.StartTime, httpParams.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(httpList, "获取成功", context)
}

func GetHttpStage(context *gin.Context) {
	var queryPageHttp request.HttpParams
	_ = context.BindQuery(&queryPageHttp)
	stageTime, err := services.GetHttpStage(queryPageHttp.MonitorId, queryPageHttp.TimeGrain, queryPageHttp.StartTime, queryPageHttp.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(stageTime, "获取成功", context)
}

func GetHttpQuota(context *gin.Context) {
	var queryPageHttp request.HttpParams
	_ = context.BindQuery(&queryPageHttp)
	quota, err := services.GetHttpQuota(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(quota, "获取成功", context)
}
