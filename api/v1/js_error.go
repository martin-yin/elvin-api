package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetJsError(context *gin.Context) {
	responses, err := services.GetJsError()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetJsErrorDetail(context *gin.Context) {
	var jsErrorParams request.JsErrorParams
	_ = context.BindQuery(&jsErrorParams)

	responses, err := services.GetJsErrorDetail(jsErrorParams.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
