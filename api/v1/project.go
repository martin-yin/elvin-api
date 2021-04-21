package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetProjectList(context *gin.Context) {
	projectList := services.GetProjectList()
	response.OkWithDetailed(projectList, "获取成功", context)
}
