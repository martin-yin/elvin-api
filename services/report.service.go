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

func CreatePageJsError(jsErrorBody *request.JsErrorBody, commonFiles *model.CommonFiles) {
	jsErrorModel := model.PageJsError{
		PageUrl:       jsErrorBody.PageUrl,
		ComponentName: jsErrorBody.ComponentName,
		Stack:         jsErrorBody.Stack,
		Message:       jsErrorBody.Message,
		StackFrames:   jsErrorBody.StackFrames,
		ErrorName:     jsErrorBody.ErrorName,
		CommonFiles:   *commonFiles,
	}
	global.GORMDB.Transaction(func(tx *gorm.DB) error {
		issueModel, err := FindIssue(jsErrorModel.Stack)
		if err == nil && issueModel.ID != 0 {
			jsErrorModel.IssuesId = issueModel.ID
			if err := tx.Create(&jsErrorModel).Error; err != nil {
				global.LOGGER.Error("创建JS Error失败！:", zap.Any("err", err))
				return err
			}
		} else {
			issueModel := model.Issue{
				ErrorName:  jsErrorBody.ErrorName,
				Message:    jsErrorBody.Message,
				MonitorId:  jsErrorBody.CommonFiles.MonitorId,
				IsFix:      false,
				FixTime:    0,
				FixUserId:  0,
				Stack:      jsErrorBody.Stack,
				HappenTime: jsErrorBody.HappenTime,
				PageJsError: []model.PageJsError{
					jsErrorModel,
				},
			}
			if err := tx.Create(&issueModel).Error; err != nil {
				global.LOGGER.Error("创建JS Error失败！:", zap.Any("err", err))
				return err
			}
		}
		return nil
	})
}

func CreatePageView(pageView *request.PageViewBody, commonFiles *model.CommonFiles) {
	pageViewModel := model.PageView{
		PageUrl:       pageView.PageUrl,
		DocumentTitle: pageView.DocumentTitle,
		Referrer:      pageView.Referrer,
		Encode:        pageView.Encode,
		CommonFiles:   *commonFiles,
	}
	global.GORMDB.Create(&pageViewModel)
}
