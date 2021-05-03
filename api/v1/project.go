package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"github.com/gin-gonic/gin"
)

func GetProjectList(context *gin.Context) {
	claims, exists := context.Get("claims")
	if exists {
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		projectList := services.GetProjectList(customClaims.ID)
		response.OkWithDetailed(projectList, "获取成功", context)
		return
	} else {
		//response.OkWithDetailed([], "获取失败", context)
	}
}
