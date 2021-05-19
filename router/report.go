package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("report")
	{
		WebPageReport.POST("performance", v1.CreatePagePerformance)
		WebPageReport.POST("http", v1.CreateHttpInfo)
		WebPageReport.POST("resource", v1.CreateResourcesError)
		WebPageReport.POST("operation", v1.CreatePageOperation)
		WebPageReport.POST("issues", v1.CreatePageIssues)
		WebPageReport.POST("view", v1.CreatePageView)
	}
}
