package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetJsError(context *gin.Context) {
	responses, err := services.GetJsError()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
