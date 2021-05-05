package router

import (
	v1 "danci-api/api/v1"
	"danci-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitCommunal(Router *gin.RouterGroup) {
	Communal := Router.Group("communal").Use(middleware.Auth())
	{
		Communal.GET("performance", v1.GetPerformance)
		Communal.GET("http", v1.GetHttpInfo)
		Communal.GET("httpStage", v1.GetHttpStage)
		Communal.GET("error", v1.GetResourceErrorInfo)
		Communal.GET("jsError", v1.GetJsError)

		Communal.GET("users", v1.GetUsers)
		Communal.GET("user", v1.GetUser)
		Communal.GET("userAction", v1.GetUserAction)
		Communal.GET("usersActionsStatistics", v1.GetUsersActionsStatistics)
		Communal.GET("surveyStatistics", v1.GetSurveyStatisticsData)
		Communal.GET("surveyPUv", v1.GetSurveyPUvData)
		Communal.GET("surveyJsError", v1.GetSurveyJsErrorData)
		Communal.GET("usersActionList", v1.GetUserActionList)
		Communal.GET("projects", v1.GetProjectList)
		Communal.GET("http-error", v1.GetHttpErrorInfo)

		// 获取团队列表
		Communal.GET("teamList", v1.GetTeamList)
		// 创建团队
		Communal.POST("createTeam", v1.CreateTeam)
		// 团队绑定管理员
		Communal.POST("bindTeamAdmins", v1.BindTeamAdmins)
		// 创建团队项目
		Communal.POST("addTeamProject", v1.AddTeamProject)

		//Communal.GET("getHttp404", func(context *gin.Context) {
		//	context.JSON(404, gin.H{"message": "hello world"})
		//})
		//Communal.GET("getHttp304", func(context *gin.Context) {
		//	context.JSON(304, gin.H{"message": "hello world"})
		//})
		//Communal.GET("getHttp500", func(context *gin.Context) {
		//	context.JSON(500, gin.H{"message": "hello world"})
		//})
		//Communal.POST("postHttp500", func(context *gin.Context) {
		//	context.JSON(500, gin.H{"message": "hello world"})
		//})
		//Communal.POST("postHttp404", func(context *gin.Context) {
		//	context.JSON(404, gin.H{"message": "hello world"})
		//})
		//Communal.POST("postHttp304", func(context *gin.Context) {
		//	context.JSON(304, gin.H{"message": "hello world"})
		//})
		//
		//Communal.POST("postHttp200", func(context *gin.Context) {
		//	context.JSON(200, gin.H{"message": "hello world"})
		//})
	}
}
