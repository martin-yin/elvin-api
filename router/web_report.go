package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPageReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("webPageReport")
	{
		WebPageReport.POST("/webReport/loadPage", v1.SetWebLoadPageInfo)
		WebPageReport.POST("/webReport/httpInfo", v1.SetWebHttpInfo)

		WebPageReport.POST("/webReport/resourceError", v1.SetWebResourcesError)
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
