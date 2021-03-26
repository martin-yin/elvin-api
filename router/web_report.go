package router

import (
	v1 "danci-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPageReport(Router *gin.RouterGroup) {
	WebPageReport := Router.Group("webPageReport")
	{
		// web页面被访问一次就会上报数据
		WebPageReport.POST("/webReport/pageView", v1.SetWebPerformance)
		// 上报请求，不管这个请求得结果是正确还是错误得！
		WebPageReport.POST("/webReport/request", v1.SetWebRequest)
		// 接受资源错误！
		WebPageReport.POST("/webReport/resourcesError", v1.SetWebResourcesError)
		// js error 这里去接受js得报错。
		WebPageReport.POST("/webReport/jsError", v1.SetWebRequest)


		// 获取资源错误pv
		WebPageReport.GET("/webReport/webResourcesError", v1.GetWebResourcesError)
	}
}
