package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"

	"github.com/gin-gonic/gin"
)

func GetIssues(context *gin.Context) {
	var params request.RequestParams
	_ = context.BindQuery(&params)
	params.StartTime += " 00:00:00"
	params.EndTime += " 23:59:59"
	responses, err := services.GetIssues(params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetJsErrorDetail(context *gin.Context) {
	var params request.JsErrorParams
	_ = context.BindQuery(&params)

	responses, err := services.GetJsErrorDetail(params.IssueId, params.ErrorId, params.MonitorId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
