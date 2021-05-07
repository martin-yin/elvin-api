package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"github.com/gin-gonic/gin"
)

func GetProjectList(context *gin.Context) {
	emptyList := make([]int, 0)
	claims, exists := context.Get("claims")
	if exists {
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		projectList, err := services.GetProjectList(customClaims.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), context)
		} else {
			if len(projectList) == 0 {
				response.OkWithDetailed(emptyList, "获取成功", context)
			} else {
				response.OkWithDetailed(projectList, "获取成功", context)
			}
		}
		return
	}
}
