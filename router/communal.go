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
		Communal.GET("error", v1.GetResourceErrorInfo)
		Communal.GET("jsError", v1.GetJsError)

		Communal.GET("users", v1.GetUsers)
		Communal.GET("user", v1.GetUser)
		Communal.GET("userAction", v1.GetUserAction)
		Communal.GET("userActions", v1.GetUserActions)
	}
}
