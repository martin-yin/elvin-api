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
	var params request.RequestParams
	err := context.BindQuery(&params)
	params.StartTime += " 00:00:00"
	params.EndTime += " 23:59:59"
	httpList, err := services.GetHttpList(params)
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
	var params request.RequestParams
	_ = context.BindQuery(&params)

	params.StartTime = params.StartTime + " 00:00:00"
	params.EndTime = params.EndTime + " 23:59:59"
	stageTime, err := services.GetHttpStage(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(stageTime, "获取成功", context)
}

func GetHttpQuota(context *gin.Context) {
	var params request.RequestParams
	_ = context.BindQuery(&params)
	params.StartTime = params.StartTime + " 00:00:00"
	params.EndTime = params.EndTime + " 23:59:59"
	quota, err := services.GetHttpQuota(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(quota, "获取成功", context)
}
