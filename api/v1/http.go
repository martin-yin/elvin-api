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
	var parmas request.RequestParams
	err := context.BindQuery(&parmas)
	parmas.StartTime = " 00:00:00"
	parmas.EndTime =  " 23:59:59"
	httpList, err := services.GetHttpInfoList(parmas)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(httpList, "获取成功", context)
}

func GetHttpErrorList(context *gin.Context) {
	var params request.RequestParams
	err := context.BindQuery(&params)
	httpList, err := services.GetHttpErrorList(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(httpList, "获取成功", context)
}

func GetHttpStage(context *gin.Context) {
	var queryPageHttp request.HttpParams
	_ = context.BindQuery(&queryPageHttp)
	stageTime, err := services.GetHttpStage(queryPageHttp.MonitorId, queryPageHttp.StartTime, queryPageHttp.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(stageTime, "获取成功", context)
}

func GetHttpQuota(context *gin.Context) {
	var params request.RequestParams
	_ = context.BindQuery(&params)
	quota, err := services.GetHttpQuota(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(quota, "获取成功", context)
}
