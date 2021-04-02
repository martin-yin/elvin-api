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
	}
}