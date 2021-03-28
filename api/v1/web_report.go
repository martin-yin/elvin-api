package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 往数据中添加web性能
func SetWebLoadPageInfo(context *gin.Context) {
	var webLoadPageInfo request.WebLoadPageInfo
	err := context.BindJSON(&webLoadPageInfo)
	if err != nil {
		fmt.Print(err)
	}
	webLoadPageInfoModel := &model.WebLoadpageInfo{
		PageUrl: webLoadPageInfo.PageUrl,
		UserId: "",
		UploadType: "LOAD_PAGE",
		HappenTime: "",
		HappenDate: "",
		PageKey: webLoadPageInfo.PageKey,

		LoadPage: webLoadPageInfo.LoadPage,
		DomReady: webLoadPageInfo.DomReady,
		Redirect: webLoadPageInfo.Redirect,

		LookupDomain: webLoadPageInfo.LookupDomain,
		Ttfb: webLoadPageInfo.Ttfb,
		Request: webLoadPageInfo.Request,
		LoadEvent: webLoadPageInfo.LoadEvent,
		Appcache: webLoadPageInfo.Appcache,
		UnloadEvent: webLoadPageInfo.UnloadEvent,

		Connect: webLoadPageInfo.Connect,
		LoadType: webLoadPageInfo.LoadType,
		BrowserInfo: webLoadPageInfo.BrowserInfo,
	}
	services.SetWebLoadPageInfo(*webLoadPageInfoModel, context)
}

// 存储页面得请求数据
func SetWebHttpInfo(context *gin.Context) {
	var webHttpInfo request.WebHttpInfo
	err := context.BindJSON(&webHttpInfo)
	if err != nil {
		fmt.Print(err)
	}

	webHttpInfoModel := &model.WebHttpInfo{
		PageUrl: webHttpInfo.PageUrl,
		UserId: "",
		UploadType: "HTTP_LOG",
		HappenTime: "",
		HappenDate: "",
		PageKey: webHttpInfo.PageKey,

		HttpUrl: webHttpInfo.HttpUrl,
		SimpleHttpUrl: webHttpInfo.SimpleHttpUrl,
		LoadTime: webHttpInfo.LoadTime,
		Status: webHttpInfo.Status,
		StatusText: webHttpInfo.StatusText,
		StatusResult: webHttpInfo.StatusResult,
		RequestText: webHttpInfo.RequestText,
		ResponseText: webHttpInfo.ResponseText,
	}
	services.WebHttpInfoModel(*webHttpInfoModel, context)
}

func SetWebResourcesError(context *gin.Context) {
	var webResourceErrorInfo request.WebResourceErrorInfo
	err := context.BindJSON(&webResourceErrorInfo)
	if err != nil {
		fmt.Print(err, "err!")
	}
	webResourceErrorInfoModel := &model.WebResourceErrorInfo{
		PageUrl: webResourceErrorInfo.PageUrl,
		UserId: webResourceErrorInfo.UserId,
		HappenTime: webResourceErrorInfo.HappenTime,
		UploadType: webResourceErrorInfo.PageUrl,
		HappenDate: webResourceErrorInfo.HappenDate,
		PageKey: webResourceErrorInfo.PageKey,
		SourceUrl: webResourceErrorInfo.SourceUrl,
		ElementType: webResourceErrorInfo.ElementType,
		Status: webResourceErrorInfo.Status,
	}
	services.SetWebResourcesError(*webResourceErrorInfoModel, context)
}

// 将uuid 写入cookie， 一个cookie 就是一个用户。每次切换一个页面等页面加载完成的时候，就是上报一次pv。
//  cookie 数量 = 用户数量
// 记录每次报错来自哪个cookie， 然后连表分析 报错影响的 用户数量 和 PV
// 计算一分钟之内的影响 =  一分钟之内有多少个报错   一分钟之内有多少个PV 一分钟之内有多少个用户
// 获取apikey 获取对应项目得资源错误，计算资源异常PV 次数 和 影响用户数量
