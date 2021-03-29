package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetWebLoadPageInfo(context *gin.Context) {
	result := services.GetWebLoadPageInfo()
	context.JSON(200, gin.H{"data": result, "message": "ok", "code": 200})
}

func SetWebLoadPageInfo(context *gin.Context) {
	var webLoadPageInfo request.WebLoadPageInfo
	err := context.BindJSON(&webLoadPageInfo)
	if err != nil {
		fmt.Print(err)
	}
	webLoadPageInfoModel := &model.WebLoadpageInfo{
		PageUrl:    webLoadPageInfo.PageUrl,
		UserId:     "",
		UploadType: "LOAD_PAGE",
		HappenTime: webLoadPageInfo.HappenTime,

		Redirect:     webLoadPageInfo.Redirect,
		Appcache:     webLoadPageInfo.Appcache,
		LookupDomain: webLoadPageInfo.LookupDomain,
		Tcp:          webLoadPageInfo.Tcp,
		SslT:         webLoadPageInfo.SslT,

		Request:  webLoadPageInfo.Request,
		DomParse: webLoadPageInfo.DomParse,

		Ttfb:      webLoadPageInfo.Ttfb,
		LoadPage:  webLoadPageInfo.LoadPage,
		LoadEvent: webLoadPageInfo.LoadEvent,

		LoadType: webLoadPageInfo.LoadType,

		DeviceName:     webLoadPageInfo.DeviceName,
		Os:             webLoadPageInfo.Os,
		BrowserName:    webLoadPageInfo.BrowserName,
		BrowserVersion: webLoadPageInfo.BrowserVersion,
		UA:             webLoadPageInfo.UA,
	}
	services.SetWebLoadPageInfo(*webLoadPageInfoModel)
}

// 存储HTTP请求
func SetWebHttpInfo(context *gin.Context) {
	var webHttpInfo []request.WebHttpInfo
	err := context.BindJSON(&webHttpInfo)
	if err != nil {
		fmt.Print(err)
	}
	var webHttpInfoModel []*model.WebHttpInfo
	for _, item := range webHttpInfo {
		webHttpInfoModelItem := &model.WebHttpInfo{
			PageUrl:      item.PageUrl,
			UserId:       "",
			UploadType:   item.UploadType,
			HappenTime:   item.HappenTime,
			HttpUrl:      item.HttpUrl,
			LoadTime:     item.LoadTime,
			Status:       item.Status,
			StatusText:   item.StatusText,
			StatusResult: item.StatusResult,
			RequestText:  item.RequestText,
			ResponseText: item.ResponseText,

			DeviceName:     item.DeviceName,
			Os:             item.Os,
			BrowserName:    item.BrowserName,
			BrowserVersion: item.BrowserVersion,
			UA:             item.UA,
		}
		webHttpInfoModel = append(webHttpInfoModel, webHttpInfoModelItem)
	}
	services.WebHttpInfoModel(webHttpInfoModel)
}

func SetWebResourcesError(context *gin.Context) {
	var webResourceErrorInfo request.WebResourceErrorInfo
	err := context.BindJSON(&webResourceErrorInfo)
	if err != nil {
		fmt.Print(err, "err!")
	}
	webResourceErrorInfoModel := &model.WebResourceErrorInfo{
		PageUrl: webResourceErrorInfo.PageUrl,
		UserId:  webResourceErrorInfo.UserId,

		HappenTime:  webResourceErrorInfo.HappenTime,
		UploadType:  webResourceErrorInfo.PageUrl,
		SourceUrl:   webResourceErrorInfo.SourceUrl,
		ElementType: webResourceErrorInfo.ElementType,
		Status:      webResourceErrorInfo.Status,

		DeviceName:     webResourceErrorInfo.DeviceName,
		Os:             webResourceErrorInfo.Os,
		BrowserName:    webResourceErrorInfo.BrowserName,
		BrowserVersion: webResourceErrorInfo.BrowserVersion,
		UA:             webResourceErrorInfo.UA,
	}
	services.SetWebResourcesError(*webResourceErrorInfoModel)
}

// 存储用户行为(点击等等……)。
func SetBehaviorInfo(context *gin.Context) {
	var webBehaviorInfo request.WebBehaviorInfo
	err := context.BindJSON(&webBehaviorInfo)
	if err != nil {
		fmt.Print(err, "err!")
	}
	webBehaviorInfoModel := &model.WebBehaviorInfo{
		PageUrl:    webBehaviorInfo.PageUrl,
		UserId:     webBehaviorInfo.UserId,
		HappenTime: webBehaviorInfo.HappenTime,
		UploadType: webBehaviorInfo.PageUrl,

		BehaviorType: webBehaviorInfo.BehaviorType,
		ClassName:    webBehaviorInfo.ClassName,
		Placeholder:  webBehaviorInfo.Placeholder,
		InputValue:   webBehaviorInfo.InputValue,
		TagNameint:   webBehaviorInfo.TagNameint,
		InnterText:   webBehaviorInfo.InnterText,

		DeviceName:     webBehaviorInfo.DeviceName,
		Os:             webBehaviorInfo.Os,
		BrowserName:    webBehaviorInfo.BrowserName,
		BrowserVersion: webBehaviorInfo.BrowserVersion,
		UA:             webBehaviorInfo.UA,
	}
	services.SetBehaviorInfo(*webBehaviorInfoModel)
}

// 将uuid 写入cookie， 一个cookie 就是一个用户。每次切换一个页面等页面加载完成的时候，就是上报一次pv。
//  cookie 数量 = 用户数量
// 记录每次报错来自哪个cookie， 然后连表分析 报错影响的 用户数量 和 PV
// 计算一分钟之内的影响 =  一分钟之内有多少个报错   一分钟之内有多少个PV 一分钟之内有多少个用户
// 获取apikey 获取对应项目得资源错误，计算资源异常PV 次数 和 影响用户数量
