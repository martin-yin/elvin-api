package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("report")
	{
		WebPageReport.POST("loadPage", v1.CreateLoadPageInfo)
		WebPageReport.POST("httpInfo", v1.CreateHttpInfo)
		WebPageReport.POST("resourceError", v1.CreateResourcesError)
		WebPageReport.POST("behavior", v1.CreateBehaviorInfo)

		WebPageReport.POST("tests", func(context *gin.Context) {
			context.JSON(200, gin.H{"result": nil, "message": "msg", "redirect_url": "url"})
		})
	}
}
