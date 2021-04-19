package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetHomeProjectStatisticsData(context *gin.Context) {
	// 获取首页的项目统计数据
	// 每小时统计一次
	//services.GetHomeProjectStatisticsData()
}

func GetPorjectList(context *gin.Context) {
	projectList := services.GetProjectList();
	response.OkWithDetailed(projectList, "获取成功", context)
}
