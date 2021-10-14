package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreatePagePerformance(performance *request.PerformanceBody, commonFiles *model.CommonFiles) {
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
		CommonFiles:  *commonFiles,
	}
	global.GORMDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&performanceModel).Error; err != nil {
			global.LOGGER.Error("写入 performance  失败！！！！！！:", zap.Any("err", err))
			return err
		}
		if err := tx.Create(&model.User{
			CommonFiles: *commonFiles,
		}).Error; err != nil {
			global.LOGGER.Error("创建用户失败！:", zap.Any("err", err))
			return err
		}
		return nil
	})
}

func CreateUserAction(commonFiles model.CommonFiles, reportData string) {
	userAction := model.UserAction{
		UserId:       commonFiles.UserId,
		MonitorId:    commonFiles.MonitorId,
		HappenTime:   commonFiles.HappenTime,
		HappenDay:    commonFiles.HappenDay,
		SessionId:    commonFiles.SessionId,
		ActionType:   commonFiles.ActionType,
		ActionDetail: reportData,
	}
	global.GORMDB.Create(&userAction)
}

func CreatePageHttp(http *request.HttpBody, commonFiles *model.CommonFiles) {
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
		CommonFiles:  *commonFiles,
	}
	global.GORMDB.Create(&httpModel)
}

func CreateResourcesError(resource *request.ResourceErrorBody, commonFiles *model.CommonFiles) {
	resourceModel := model.PageResourceError{
		PageUrl:     resource.PageUrl,
		SourceUrl:   resource.SourceUrl,
		ElementType: resource.ElementType,
		Status:      resource.Status,
		CommonFiles: *commonFiles,
	}
	global.GORMDB.Create(&resourceModel)
}

func CreatePageOperation(operation *request.OperationBody, commonFiles *model.CommonFiles) {
	operationModel := model.PageOperation{
		PageUrl:     operation.PageUrl,
		ClassName:   operation.ClassName,
		Placeholder: operation.Placeholder,
		InputValue:  operation.InputValue,
		TagName:     operation.TagName,
		InnerText:   operation.InnerText,
		CommonFiles: *commonFiles,
	}
	global.GORMDB.Create(&operationModel)
}

func CreatePageJsError(issue *request.IssuesBody, commonFiles *model.CommonFiles) {
	issueModel := model.PageIssue{
		PageUrl:       issue.PageUrl,
		ComponentName: issue.ComponentName,
		Stack:         issue.Stack,
		Message:       issue.Message,
		StackFrames:   issue.StackFrames,
		ErrorName:     issue.ErrorName,
		CommonFiles:   *commonFiles,
	}
	jsIssueModel, err := FindJsIssue(issueModel.Message)
	if err == nil {
		if jsIssueModel.ID != 0 {
			issueModel.IssuesId = jsIssueModel.ID
			global.GORMDB.Save(&issueModel)
		} else {
			jsIssue := model.Issue{
				ErrorName:  issue.ErrorName,
				Message:    issue.Message,
				MonitorId:  issue.CommonFiles.MonitorId,
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
	global.GORMDB.Save(&stack)
}

func CreatePageView(pageView *request.PageViewBody, commonFiles *model.CommonFiles) {
	pageViewModel := model.PageView{
		PageUrl:     pageView.PageUrl,
		CommonFiles: *commonFiles,
	}
	global.GORMDB.Create(&pageViewModel)
}
