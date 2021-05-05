package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAdmin(Router *gin.RouterGroup) {
	Admin := Router.Group("admin")
	{
		Admin.POST("createAdmin", v1.CreateAdmin)
		Admin.POST("adminLogin", v1.AdminLogin)
		Admin.GET("teamList", v1.GetTeamList)
		Admin.POST("createTeam", v1.CreateTeam)
	}
}
