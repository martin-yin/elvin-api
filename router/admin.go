package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAdmin(Router *gin.RouterGroup) {
	Admin := Router.Group("admin")
	{
		Admin.POST("registerAdmin", v1.RegisterAdmin)
		Admin.POST("adminLogin", v1.AdminLogin)
	}
}
