package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"fmt"
)

func CreatePagePerformance(performance *request.PerformanceBody, publicFiles *model.PublicFiles) {
	performanceModel := &model.PagePerformance{
		PageUrl:      performance.PageUrl,
		Appcache:     performance.Appcache,
		LookupDomain: performance.LookupDomain,
		Tcp:          performance.Tcp,
		SslT:         performance.SslT,
		Request:      performance.Request,
		DomParse:     performance.DomParse,
		Ttfb:         performance.Ttfb,
		LoadPage:     performance.LoadPage,
		LoadEvent:    performance.LoadEvent,
		LoadType:     performance.LoadType,
		Redirect:     performance.Redirect,
		PublicFiles:  *publicFiles,
	}
	if err := global.GVA_DB.Create(&performanceModel).Error; err != nil {
		fmt.Println("err", err)
	}
	user := model.User{
		PublicFiles: performance.PublicFiles,
	}
	CreateUser(&user)
}

func CreateUser(user *model.User) {
	// 用户每次刷新页面就记录一次用户。
	if err := global.GVA_DB.Create(&user).Error; err != nil {
		fmt.Println(err, "创建用户出错！")
	}
}

func CreateUserAction(publicFiles model.PublicFiles, reportData string) {
	userAction := model.UserAction{
		UserId:       publicFiles.UserId,
		MonitorId:    publicFiles.MonitorId,
		HappenTime:   publicFiles.HappenTime,
		HappenDay:    publicFiles.HappenDay,
		EventId:      publicFiles.EventId,
		ActionType:   publicFiles.ActionType,
		ActionDetail: reportData,
	}
	global.GVA_DB.Create(&userAction)
}

func CreatePageHttp(http *request.HttpBody, publicFiles *model.PublicFiles) {
	httpModel := model.PageHttp{
		PageUrl:      http.PageUrl,
		HttpUrl:      http.HttpUrl,
		LoadTime:     http.LoadTime,
		Method:       http.Method,
		Status:       http.Status,
		StatusText:   http.StatusText,
		StatusResult: http.StatusResult,
		RequestText:  http.RequestText,
		ResponseText: http.ResponseText,
		PublicFiles:  *publicFiles,
	}
	global.GVA_DB.Create(&httpModel)
}

func CreateResourcesError(resource *request.ResourceErrorBody, publicFiles *model.PublicFiles) {
	resourceModel := model.PageResourceError{
		PageUrl:     resource.PageUrl,
		SourceUrl:   resource.SourceUrl,
		ElementType: resource.ElementType,
		Status:      resource.Status,
		PublicFiles: *publicFiles,
	}
	global.GVA_DB.Create(&resourceModel)
}

func CreatePageOperation(operation *request.OperationBody, publicFiles *model.PublicFiles) {
	operationModel := model.PageOperation{
		PageUrl:     operation.PageUrl,
		ClassName:   operation.ClassName,
		Placeholder: operation.Placeholder,
		InputValue:  operation.InputValue,
		TagName:     operation.TagName,
		InnerText:   operation.InnerText,
		PublicFiles: *publicFiles,
	}
	global.GVA_DB.Create(&operationModel)
}

func CreatePageJsError(issue *request.IssuesBody, publicFiles *model.PublicFiles) {
	issueModel := model.PageIssue{
		PageUrl:       issue.PageUrl,
		ComponentName: issue.ComponentName,
		Stack:         issue.Stack,
		Message:       issue.Message,
		StackFrames:   issue.StackFrames,
		ErrorName:     issue.ErrorName,
		PublicFiles:   *publicFiles,
	}
	jsIssueModel, err := FindJsIssue(issueModel.Message)
	if err == nil {
		if jsIssueModel.ID != 0 {
			issueModel.IssuesId = jsIssueModel.ID
			global.GVA_DB.Save(&issueModel)
		} else {
			jsIssue := model.Issue{
				ErrorName: issue.ErrorName,
				Message:   issue.Message,
				MonitorId: issue.PublicFiles.MonitorId,
				HappenTime: issue.HappenTime,
				PageIssue: []model.PageIssue{
					issueModel,
				},
			}
			CreateJsIssue(jsIssue)
		}
	}
}

func CreateJsIssue(stack model.Issue) {
	global.GVA_DB.Save(&stack)
}

func CreatePageView(pageView *request.PageViewBody, publicFiles *model.PublicFiles) {
	pageViewModel := model.PageView{
		PageUrl:     pageView.PageUrl,
		PublicFiles: *publicFiles,
	}
	global.GVA_DB.Create(&pageViewModel)
}
