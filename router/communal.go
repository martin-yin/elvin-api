package router

import (
	v1 "dancin-api/api/v1"
	"dancin-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommunal(Router *gin.RouterGroup) {
	Communal := Router.Group("communal")
	{
		//performance 相关API接口
		Communal.GET("performanceStack", v1.GetPerformanceStack)
		Communal.GET("performancePages", v1.GetPerformancePageList)
		Communal.GET("performanceStageTime", v1.GetPerformanceStageTime)
		//Communal.GET("performanceRankingList", v1.GetPerformanceRankingList)
		Communal.GET("performanceQuota", v1.GetPerformanceQuota)

		// http相关API接口
		Communal.GET("https", v1.GetHttpList)
		Communal.GET("httpStage", v1.GetHttpStage)
		Communal.GET("httpQuota", v1.GetHttpQuota)
		Communal.GET("httpErrors", v1.GetHttpErrorList)

		// 用户相关接口
		Communal.GET("users", v1.GetUsers)
		Communal.GET("user", v1.GetUser)
		Communal.GET("userAction", v1.GetUserAction)
		Communal.GET("userActionStatistics", v1.GetUsersActionsStatistics)
		Communal.GET("userActions", v1.GetUserActionList)

		// 项目健康状态
		Communal.Use(middleware.Auth()).GET("getHealthStatus", v1.GetHealthStatus)
		// 资源错误接口
		Communal.GET("staticErr", v1.GetResourceError)
		Communal.GET("jsErrors", v1.GetIssues)
		Communal.GET("jsError", v1.GetIssuesDetail)
	}
}
