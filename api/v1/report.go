package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func CreatePagePerformance(context *gin.Context) {
	var pagePerformanceBody request.PostPagePerformance
	err := context.BindJSON(&pagePerformanceBody)
	pagePerformanceModel := model.PagePerformance{
		PageUrl:        pagePerformanceBody.PageUrl,
		UserId:         pagePerformanceBody.UserId,
		ApiKey:         pagePerformanceBody.ApiKey,
		ActionType:     pagePerformanceBody.ActionType,
		HappenTime:     pagePerformanceBody.HappenTime,
		HappenDay:      pagePerformanceBody.HappenDay,
		Redirect:       pagePerformanceBody.Redirect,
		Appcache:       pagePerformanceBody.Appcache,
		LookupDomain:   pagePerformanceBody.LookupDomain,
		Tcp:            pagePerformanceBody.Tcp,
		SslT:           pagePerformanceBody.SslT,
		Request:        pagePerformanceBody.Request,
		DomParse:       pagePerformanceBody.DomParse,
		Ttfb:           pagePerformanceBody.Ttfb,
		LoadPage:       pagePerformanceBody.LoadPage,
		LoadEvent:      pagePerformanceBody.LoadEvent,
		LoadType:       pagePerformanceBody.LoadType,
		IP:             context.ClientIP(),
		Device:         pagePerformanceBody.Device,
		DeviceType:     pagePerformanceBody.DeviceType,
		Os:             pagePerformanceBody.Os,
		OsVersion:      pagePerformanceBody.OsVersion,
		Browser:        pagePerformanceBody.Browser,
		BrowserVersion: pagePerformanceBody.BrowserVersion,
		UA:             pagePerformanceBody.UA,
	}

	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	if err := services.CreatePagePerformance(pagePerformanceModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

// 存储HTTP请求
func CreateHttpInfo(context *gin.Context) {
	var pageHttpBody request.PostPageHttpBody
	_ = context.BindJSON(&pageHttpBody)
	webHttpInfoModel := model.PageHttp{
		PageUrl:      pageHttpBody.PageUrl,
		UserId:       pageHttpBody.UserId,
		ApiKey:       pageHttpBody.ApiKey,
		ActionType:   pageHttpBody.ActionType,
		HappenTime:   pageHttpBody.HappenTime,
		HappenDay:    pageHttpBody.HappenDay,
		HttpUrl:      pageHttpBody.HttpUrl,
		LoadTime:     pageHttpBody.LoadTime,
		Status:       pageHttpBody.Status,
		StatusText:   pageHttpBody.StatusText,
		StatusResult: pageHttpBody.StatusResult,
		RequestText:  pageHttpBody.RequestText,
		ResponseText: pageHttpBody.ResponseText,

		IP:             context.ClientIP(),
		Device:         pageHttpBody.Device,
		DeviceType:     pageHttpBody.DeviceType,
		Os:             pageHttpBody.Os,
		OsVersion:      pageHttpBody.OsVersion,
		Browser:        pageHttpBody.Browser,
		BrowserVersion: pageHttpBody.BrowserVersion,
		UA:             pageHttpBody.UA,
	}

	if err := services.CreatePageHttpModel(webHttpInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreateResourcesError(context *gin.Context) {
	var pageResourceErroBody request.PostPageResourceErroBody
	_ = context.BindJSON(&pageResourceErroBody)

	resourceErrorInfoModel := model.PageResourceError{
		PageUrl:        pageResourceErroBody.PageUrl,
		UserId:         pageResourceErroBody.UserId,
		ApiKey:         pageResourceErroBody.ApiKey,
		HappenTime:     pageResourceErroBody.HappenTime,
		HappenDay:      pageResourceErroBody.HappenDay,
		ActionType:     pageResourceErroBody.ActionType,
		SourceUrl:      pageResourceErroBody.SourceUrl,
		ElementType:    pageResourceErroBody.ElementType,
		Status:         pageResourceErroBody.Status,
		IP:             context.ClientIP(),
		Device:         pageResourceErroBody.Device,
		DeviceType:     pageResourceErroBody.DeviceType,
		Os:             pageResourceErroBody.Os,
		OsVersion:      pageResourceErroBody.OsVersion,
		Browser:        pageResourceErroBody.Browser,
		BrowserVersion: pageResourceErroBody.BrowserVersion,
		UA:             pageResourceErroBody.UA,
	}

	if err := services.CreateResourcesError(resourceErrorInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreatePageBehavior(context *gin.Context) {
	var behaviorInfoBody request.PostBehaviorInfoBody
	_ = context.BindJSON(&behaviorInfoBody)

	pageBehaviorInfoModel := model.PageBehavior{
		PageUrl:     behaviorInfoBody.PageUrl,
		UserId:      behaviorInfoBody.UserId,
		ApiKey:      behaviorInfoBody.ApiKey,
		HappenTime:  behaviorInfoBody.HappenTime,
		HappenDay:   behaviorInfoBody.HappenDay,
		ActionType:  behaviorInfoBody.ActionType,
		ClassName:   behaviorInfoBody.ClassName,
		Placeholder: behaviorInfoBody.Placeholder,
		InputValue:  behaviorInfoBody.InputValue,
		TagNameint:  behaviorInfoBody.TagNameint,
		InnterText:  behaviorInfoBody.InnterText,

		IP:             context.ClientIP(),
		Device:         behaviorInfoBody.Device,
		DeviceType:     behaviorInfoBody.DeviceType,
		Os:             behaviorInfoBody.Os,
		OsVersion:      behaviorInfoBody.OsVersion,
		Browser:        behaviorInfoBody.Browser,
		BrowserVersion: behaviorInfoBody.BrowserVersion,
		UA:             behaviorInfoBody.UA,
	}

	if err := services.CreatePageBehavior(pageBehaviorInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreatePageJsError(context *gin.Context) {
	var jsErrorInfoBody request.PostJsErrorInfoBody
	_ = context.BindJSON(&jsErrorInfoBody)
	jsErrorInfoModel := model.PageJsError{
		PageUrl:        jsErrorInfoBody.PageUrl,
		UserId:         jsErrorInfoBody.UserId,
		ApiKey:         jsErrorInfoBody.ApiKey,
		HappenTime:     jsErrorInfoBody.HappenTime,
		HappenDay:      jsErrorInfoBody.HappenDay,
		ActionType:     jsErrorInfoBody.ActionType,
		ComponentName:  jsErrorInfoBody.ComponentName,
		Stack:          jsErrorInfoBody.Stack,
		Message:        jsErrorInfoBody.Message,
		IP:             context.ClientIP(),
		Device:         jsErrorInfoBody.Device,
		DeviceType:     jsErrorInfoBody.DeviceType,
		Os:             jsErrorInfoBody.Os,
		OsVersion:      jsErrorInfoBody.OsVersion,
		Browser:        jsErrorInfoBody.Browser,
		BrowserVersion: jsErrorInfoBody.BrowserVersion,
		UA:             jsErrorInfoBody.UA,
	}
	if err := services.CreatePageJsError(jsErrorInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}
