package router

import (
	v1 "danci-api/api/v1"
	"danci-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitCommunal(Router *gin.RouterGroup) {
	Communal := Router.Group("communal").Use(middleware.Auth())
	{
		//performance 相关API接口
		Communal.GET("performanceStack", v1.GetPerformanceStack)
		Communal.GET("performancePageList", v1.GetPerformancePageList)
		Communal.GET("performanceStageTime", v1.GetPerformanceStageTime)
		Communal.GET("performanceRankingList", v1.GetPerformanceRankingList)
		Communal.GET("performanceQuota", v1.GetPerformanceQuota)

		// http相关API接口
		Communal.GET("httpList", v1.GetHttpList)
		Communal.GET("httpStage", v1.GetHttpStage)
		Communal.GET("httpQuota", v1.GetHttpQuota)

		// 用户相关接口
		Communal.GET("userList", v1.GetUserList)
		Communal.GET("user", v1.GetUser)
		Communal.GET("userAction", v1.GetUserAction)
		Communal.GET("usersActionsStatistics", v1.GetUsersActionsStatistics)
		Communal.GET("usersActionList", v1.GetUserActionList)

		// survey 概况数据
		Communal.GET("surveyStatistics", v1.GetSurveyStatisticsData)
		Communal.GET("surveyPUv", v1.GetSurveyPUvData)
		Communal.GET("surveyJsError", v1.GetSurveyJsErrorData)

		//Communal.GET("projects", v1.GetProjectList)

	}
}
