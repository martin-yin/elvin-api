package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetIssues(context *gin.Context) {
	responses, err := services.GetIssues()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetIssuesDetail(context *gin.Context) {
	var jsErrorParams request.JsErrorParams
	_ = context.BindQuery(&jsErrorParams)

	responses, err := services.GetJsErrorDetail(jsErrorParams.IssueId, jsErrorParams.ErrorId, jsErrorParams.MonitorId)
	//services.GetJsErrorPreAndNext(20)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
