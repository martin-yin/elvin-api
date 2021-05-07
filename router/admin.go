package router

import (
	v1 "danci-api/api/v1"
	"danci-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdmin(Router *gin.RouterGroup) {
	Admin := Router.Group("admin")
	{
		Admin.POST("registerAdmin", v1.RegisterAdmin)
		Admin.POST("adminLogin", v1.AdminLogin)

		// 获取团队列表
		Admin.Use(middleware.Auth()).GET("teamList", v1.GetTeamList)
		// 创建团队
		Admin.Use(middleware.Auth()).POST("createTeam", v1.CreateTeam)
		// 团队绑定管理员
		Admin.Use(middleware.Auth()).POST("bindTeamAdmins", v1.BindTeamAdmins)
		// 创建团队项目
		Admin.Use(middleware.Auth()).POST("addTeamProject", v1.AddTeamProject)
		// 获取项目列表，根据当前登录的id查找项目
		Admin.Use(middleware.Auth()).GET("projectList", v1.GetProjectList)
	}
}
