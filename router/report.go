package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPageReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("report")
	{
		WebPageReport.POST("/loadPage", v1.CreateLoadPageInfo)
		WebPageReport.POST("/httpInfo", v1.CreateHttpInfo)
		WebPageReport.POST("/resourceError", v1.CreateResourcesError)
		WebPageReport.POST("/behavior", v1.CreateBehaviorInfo)
		WebPageReport.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{"result": nil, "message": "msg", "redirect_url": "url"})
		})
		WebPageReport.GET("report", v1.GetWebLoadPageInfo)
		WebPageReport.GET("http", v1.GetWebHttpInfo)
		WebPageReport.GET("error", v1.GetWebResourceErrorInfo)
		WebPageReport.POST("/tests", func(context *gin.Context) {
			context.JSON(200, gin.H{"result": nil, "message": "msg", "redirect_url": "url"})
		})
		//// 上报请求，不管这个请求得结果是正确还是错误得！
		//WebPageReport.POST("/webReport/request", v1.SetWebRequest)
		//// 接受资源错误！
		//
		//// js error 这里去接受js得报错。
		//WebPageReport.POST("/webReport/jsError", v1.SetWebRequest)
		//
		//// 获取资源错误pv
		//WebPageReport.GET("/webReport/webResourcesError", v1.GetWebResourcesError)
	}
}
