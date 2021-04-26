package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCommunal(Router *gin.RouterGroup) {
	Communal := Router.Group("communal")
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
		Communal.GET("usersActionList", v1.GetUserActionList)
		Communal.GET("projects", v1.GetProjectList)

		Communal.GET("http-error", v1.GetHttpErrorInfo)

		//Communal.GET("projectList", v1.GetUserActions)
		// 先去判断是否登录？
		// 如果没有登陆的话，先去登录，然后在登录中判断这个admin id 是否存在项目，如果没有项目的话，提示下no_project，
		// 创建项目之后 返回首页。提示下他咋接入项目等等……

		Communal.GET("getHttp404", func(context *gin.Context) {
			context.JSON(404, gin.H{"message": "hello world"})
		})
		Communal.GET("getHttp304", func(context *gin.Context) {
			context.JSON(304, gin.H{"message": "hello world"})
		})
		Communal.GET("getHttp500", func(context *gin.Context) {
			context.JSON(500, gin.H{"message": "hello world"})
		})
		Communal.POST("postHttp500", func(context *gin.Context) {
			context.JSON(500, gin.H{"message": "hello world"})
		})
		Communal.POST("postHttp404", func(context *gin.Context) {
			context.JSON(404, gin.H{"message": "hello world"})
		})
		Communal.POST("postHttp304", func(context *gin.Context) {
			context.JSON(304, gin.H{"message": "hello world"})
		})

		Communal.POST("postHttp200", func(context *gin.Context) {
			context.JSON(200, gin.H{"message": "hello world"})
		})
	}
}
