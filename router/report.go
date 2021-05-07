package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("report")
	{
		WebPageReport.POST("performance", v1.CreatePagePerformance)
		WebPageReport.POST("httpInfo", v1.CreateHttpInfo)
		WebPageReport.POST("resourceError", v1.CreateResourcesError)
		WebPageReport.POST("operation", v1.CreatePageOperation)
		WebPageReport.POST("jsError", v1.CreatePageJsError)
		WebPageReport.POST("pageView", v1.CreatePageView)
	}
}
