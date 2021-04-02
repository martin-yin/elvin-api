package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetWebLoadPageInfo(context *gin.Context) {
	result := services.GetWebLoadPageInfo()
	response.OkWithDetailed(result, "获取成功", context)
}

func GetWebHttpInfo(context *gin.Context) {
	result := services.GetWebHttpInfo()
	response.OkWithDetailed(result, "获取成功", context)
}

func GetWebResourceErrorInfo(context *gin.Context) {
	result := services.GetWebResourceErrorInfo()
	response.OkWithDetailed(result, "获取成功", context)
}

func CreateLoadPageInfo(context *gin.Context) {
	var loadPageInfoBody request.PostLoadPageInfoBody
	_ = context.BindJSON(&loadPageInfoBody)
	webLoadPageInfoModel := model.LoadpageInfo{
		PageUrl:      loadPageInfoBody.PageUrl,
		UserId:       loadPageInfoBody.UserId,
		ApiKey:       loadPageInfoBody.ApiKey,
		UploadType:   loadPageInfoBody.UploadType,
		HappenTime:   loadPageInfoBody.HappenTime,
		Redirect:     loadPageInfoBody.Redirect,
		Appcache:     loadPageInfoBody.Appcache,
		LookupDomain: loadPageInfoBody.LookupDomain,
		Tcp:          loadPageInfoBody.Tcp,
		SslT:         loadPageInfoBody.SslT,
		Request:      loadPageInfoBody.Request,
		DomParse:     loadPageInfoBody.DomParse,
		Ttfb:         loadPageInfoBody.Ttfb,
		LoadPage:     loadPageInfoBody.LoadPage,
		LoadEvent:    loadPageInfoBody.LoadEvent,
		LoadType:     loadPageInfoBody.LoadType,

		Os:             loadPageInfoBody.Os,
		OsVersion:      loadPageInfoBody.OsVersion,
		Browser:        loadPageInfoBody.Browser,
		BrowserVersion: loadPageInfoBody.BrowserVersion,
		UA:             loadPageInfoBody.UA,
	}
	if err := services.CreateLoadPageInfo(webLoadPageInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

// 存储HTTP请求
func CreateHttpInfo(context *gin.Context) {
	var httpInfoBody []request.PostHttpInfoBody
	_ = context.BindJSON(&httpInfoBody)
	var webHttpInfoModel []model.HttpInfo
	for _, item := range httpInfoBody {
		webHttpInfoModelItem := model.HttpInfo{
			PageUrl:      item.PageUrl,
			UserId:       item.UserId,
			ApiKey:       item.ApiKey,
			UploadType:   item.UploadType,
			HappenTime:   item.HappenTime,
			HttpUrl:      item.HttpUrl,
			LoadTime:     item.LoadTime,
			Status:       item.Status,
			StatusText:   item.StatusText,
			StatusResult: item.StatusResult,
			RequestText:  item.RequestText,
			ResponseText: item.ResponseText,

			Os:             item.Os,
			OsVersion:      item.OsVersion,
			Browser:        item.Browser,
			BrowserVersion: item.BrowserVersion,
			UA:             item.UA,
		}
		webHttpInfoModel = append(webHttpInfoModel, webHttpInfoModelItem)
	}

	if err := services.CreateHttpInfoModel(webHttpInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreateResourcesError(context *gin.Context) {
	var resourceErrorInfoBody request.PostResourceErrorInfoBody
	_ = context.BindJSON(&resourceErrorInfoBody)

	webResourceErrorInfoModel := model.ResourceErrorInfo{
		PageUrl:     resourceErrorInfoBody.PageUrl,
		UserId:      resourceErrorInfoBody.UserId,
		ApiKey:      resourceErrorInfoBody.ApiKey,
		HappenTime:  resourceErrorInfoBody.HappenTime,
		UploadType:  resourceErrorInfoBody.UploadType,
		SourceUrl:   resourceErrorInfoBody.SourceUrl,
		ElementType: resourceErrorInfoBody.ElementType,
		Status:      resourceErrorInfoBody.Status,

		Os:             resourceErrorInfoBody.Os,
		OsVersion:      resourceErrorInfoBody.OsVersion,
		Browser:        resourceErrorInfoBody.Browser,
		BrowserVersion: resourceErrorInfoBody.BrowserVersion,
		UA:             resourceErrorInfoBody.UA,
	}

	if err := services.CreateResourcesError(webResourceErrorInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

// 存储用户行为(点击等等……)。
func CreateBehaviorInfo(context *gin.Context) {
	var behaviorInfoBody request.PostBehaviorInfoBody
	err := context.BindJSON(&behaviorInfoBody)
	if err != nil {
		fmt.Print(err, "err!")
	}
	webBehaviorInfoModel := model.BehaviorInfo{
		PageUrl:      behaviorInfoBody.PageUrl,
		UserId:       behaviorInfoBody.UserId,
		ApiKey:       behaviorInfoBody.ApiKey,
		HappenTime:   behaviorInfoBody.HappenTime,
		UploadType:   behaviorInfoBody.UploadType,
		BehaviorType: behaviorInfoBody.BehaviorType,
		ClassName:    behaviorInfoBody.ClassName,
		Placeholder:  behaviorInfoBody.Placeholder,
		InputValue:   behaviorInfoBody.InputValue,
		TagNameint:   behaviorInfoBody.TagNameint,
		InnterText:   behaviorInfoBody.InnterText,

		Os:             behaviorInfoBody.Os,
		OsVersion:      behaviorInfoBody.OsVersion,
		Browser:        behaviorInfoBody.Browser,
		BrowserVersion: behaviorInfoBody.BrowserVersion,
		UA:             behaviorInfoBody.UA,
	}

	if err := services.CreateBehaviorInfo(webBehaviorInfoModel); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}
