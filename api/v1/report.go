package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"github.com/gin-gonic/gin"
)

func CreatePagePerformance(context *gin.Context) {
	var pagePerformanceBody request.PostPagePerformance
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&pagePerformanceBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	pagePerformanceModel := model.PagePerformance{
		PageUrl:      pagePerformanceBody.PageUrl,
		Appcache:     pagePerformanceBody.Appcache,
		LookupDomain: pagePerformanceBody.LookupDomain,
		Tcp:          pagePerformanceBody.Tcp,
		SslT:         pagePerformanceBody.SslT,
		Request:      pagePerformanceBody.Request,
		DomParse:     pagePerformanceBody.DomParse,
		Ttfb:         pagePerformanceBody.Ttfb,
		LoadPage:     pagePerformanceBody.LoadPage,
		LoadEvent:    pagePerformanceBody.LoadEvent,
		LoadType:     pagePerformanceBody.LoadType,
		Redirect:     pagePerformanceBody.Redirect,
		PublicFiles:  publicFiles,
	}
	if err := services.CreatePagePerformance(pagePerformanceModel, pagePerformanceBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

// 存储HTTP请求
func CreateHttpInfo(context *gin.Context) {
	var pageHttpBody request.PostPageHttpBody
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&pageHttpBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	webHttpInfoModel := model.PageHttp{
		PageUrl:      pageHttpBody.PageUrl,
		HttpUrl:      pageHttpBody.HttpUrl,
		LoadTime:     pageHttpBody.LoadTime,
		Status:       pageHttpBody.Status,
		StatusText:   pageHttpBody.StatusText,
		StatusResult: pageHttpBody.StatusResult,
		RequestText:  pageHttpBody.RequestText,
		ResponseText: pageHttpBody.ResponseText,
		PublicFiles:  publicFiles,
	}

	if err := services.CreatePageHttpModel(webHttpInfoModel, pageHttpBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreateResourcesError(context *gin.Context) {
	var pageResourceErroBody request.PostPageResourceErroBody
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&pageResourceErroBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	resourceErrorInfoModel := model.PageResourceError{
		PageUrl:     pageResourceErroBody.PageUrl,
		SourceUrl:   pageResourceErroBody.SourceUrl,
		ElementType: pageResourceErroBody.ElementType,
		Status:      pageResourceErroBody.Status,
		PublicFiles: publicFiles,
	}
	if err := services.CreateResourcesError(resourceErrorInfoModel, pageResourceErroBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreatePageBehavior(context *gin.Context) {
	var behaviorInfoBody request.PostBehaviorInfoBody
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&behaviorInfoBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	pageBehaviorInfoModel := model.PageBehavior{
		PageUrl:     behaviorInfoBody.PageUrl,
		ClassName:   behaviorInfoBody.ClassName,
		Placeholder: behaviorInfoBody.Placeholder,
		InputValue:  behaviorInfoBody.InputValue,
		TagNameint:  behaviorInfoBody.TagNameint,
		InnterText:  behaviorInfoBody.InnterText,
		PublicFiles: publicFiles,
	}
	if err := services.CreatePageBehavior(pageBehaviorInfoModel, behaviorInfoBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreatePageJsError(context *gin.Context) {
	var jsErrorInfoBody request.PostJsErrorInfoBody
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&jsErrorInfoBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	jsErrorInfoModel := model.PageJsError{
		PageUrl:       jsErrorInfoBody.PageUrl,
		ComponentName: jsErrorInfoBody.ComponentName,
		Stack:         jsErrorInfoBody.Stack,
		Message:       jsErrorInfoBody.Message,
		PublicFiles:   publicFiles,
	}
	if err := services.CreatePageJsError(jsErrorInfoModel, jsErrorInfoBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreatePageView(context *gin.Context) {
	var pageViewBody request.PostPageViewBody
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&pageViewBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	pageViewModel := model.PageView{
		PageUrl:     pageViewBody.PageUrl,
		PublicFiles: publicFiles,
	}
	if err := services.CreatePageView(pageViewModel, pageViewBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}
