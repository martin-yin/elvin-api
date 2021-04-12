package router

import (
	v1 "danci-api/api/v1"
	"danci-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("report").Use(middleware.PublicFields())
	{
		WebPageReport.POST("performance", v1.CreatePagePerformance)
		WebPageReport.POST("httpInfo", v1.CreateHttpInfo)
		WebPageReport.POST("resourceError", v1.CreateResourcesError)
		WebPageReport.POST("behavior", v1.CreatePageBehavior)
		WebPageReport.POST("jsError", v1.CreatePageJsError)
		WebPageReport.POST("pageView", v1.CreatePageView)
	}
}
