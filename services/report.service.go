package services

import (
	"danci-api/global"
	"danci-api/model"
	"reflect"
)

func CreatePagePerformance(pagePerformance *model.PagePerformance, eventId string) error {
	if err := global.GVA_DB.Create(&pagePerformance).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:    pagePerformance.PageUrl,
		UserId:     pagePerformance.PublicFiles.UserId,
		ApiKey:     pagePerformance.PublicFiles.ApiKey,
		HappenTime: pagePerformance.PublicFiles.HappenTime,
		HappenDay:  pagePerformance.PublicFiles.HappenDay,
		ActionType: pagePerformance.PublicFiles.ActionType,
		LoadType:   pagePerformance.LoadType,
		EventId:    eventId,
		ActionID:   pagePerformance.ID,
	}

	userModel := model.User{
		UserId:         pagePerformance.PublicFiles.UserId,
		ApiKey:         pagePerformance.PublicFiles.ApiKey,
		HappenTime:     pagePerformance.PublicFiles.HappenTime,
		HappenDay:      pagePerformance.PublicFiles.HappenDay,
		EventId:        eventId,
		IP:             pagePerformance.PublicFiles.IP,
		Device:         pagePerformance.PublicFiles.Device,
		DeviceType:     pagePerformance.PublicFiles.DeviceType,
		Os:             pagePerformance.PublicFiles.Os,
		OsVersion:      pagePerformance.PublicFiles.OsVersion,
		Browser:        pagePerformance.PublicFiles.Browser,
		BrowserVersion: pagePerformance.PublicFiles.BrowserVersion,
		UA:             pagePerformance.PublicFiles.UA,
		Nation:         pagePerformance.PublicFiles.Nation,
		Province:       pagePerformance.PublicFiles.Province,
		City:           pagePerformance.PublicFiles.City,
		District:       pagePerformance.PublicFiles.District,
	}
	err := CreateUser(userModel)
	err = CreateUserAction(userActionModel)
	return err
}

func CreateUser(user model.User) error {
	// 用户每次刷新页面就记录一次用户。
	if err := global.GVA_DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserAction(userAction model.UserAction) error {
	if err := global.GVA_DB.Create(&userAction).Error; err != nil {
		return err
	}
	return nil
}

// 存 http请求信息。
func CreatePageHttpModel(pageHttp model.PageHttp, eventId string) error {
	var pageHttpStatistical model.PageHttpStatistical
	err := global.GVA_DB.Model(&model.PageHttpStatistical{}).Where("http_url = ? And happen_day = ? ", pageHttp.HttpUrl, pageHttp.PublicFiles.HappenDay).Find(&pageHttpStatistical).Error
	if !reflect.DeepEqual(pageHttpStatistical, model.PageHttpStatistical{}) {
		pageHttpStatistical.Total++
		if pageHttp.Status > 304 {
			pageHttpStatistical.FailTotal++
		} else {
			pageHttpStatistical.SuccessTotal++
		}
	} else {
		pageHttpStatistical.PageUrl = pageHttp.PageUrl
		pageHttpStatistical.HttpUrl = pageHttp.HttpUrl
		pageHttpStatistical.Total++
		if pageHttp.Status > 304 {
			pageHttpStatistical.FailTotal++
		} else {
			pageHttpStatistical.SuccessTotal++
		}
	}
	pageHttpStatistical.HappenDay = pageHttp.PublicFiles.HappenDay;
	err = global.GVA_DB.Save(&pageHttpStatistical).Error
	err = global.GVA_DB.Create(&pageHttp).Error

	userActionModel := model.UserAction{
		PageUrl:    pageHttp.PageUrl,
		UserId:     pageHttp.PublicFiles.UserId,
		ApiKey:     pageHttp.PublicFiles.ApiKey,
		EventId:    eventId,
		HappenTime: pageHttp.PublicFiles.HappenTime,
		HappenDay:  pageHttp.PublicFiles.HappenDay,
		ActionType: pageHttp.PublicFiles.ActionType,
		ActionID:   pageHttp.ID,
		HttpUrl:    pageHttp.HttpUrl,
	}
	CreateUserAction(userActionModel)
	if err != nil {
		return err
	}
	return nil
}

func CreateResourcesError(resourceErrorInfo model.PageResourceError, eventId string) error {
	if err := global.GVA_DB.Create(&resourceErrorInfo).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:     resourceErrorInfo.PageUrl,
		UserId:      resourceErrorInfo.PublicFiles.UserId,
		ApiKey:      resourceErrorInfo.PublicFiles.ApiKey,
		HappenTime:  resourceErrorInfo.PublicFiles.HappenTime,
		HappenDay:   resourceErrorInfo.PublicFiles.HappenDay,
		ActionType:  resourceErrorInfo.PublicFiles.ActionType,
		ActionID:    resourceErrorInfo.ID,
		EventId:     eventId,
		SourceUrl:   resourceErrorInfo.SourceUrl,
		ElementType: resourceErrorInfo.ElementType,
	}
	err := CreateUserAction(userActionModel)
	return err
}

func CreatePageBehavior(pageBehavior model.PageBehavior, eventId string) error {
	err := global.GVA_DB.Create(&pageBehavior).Error
	userActionModel := model.UserAction{
		PageUrl:    pageBehavior.PageUrl,
		UserId:     pageBehavior.PublicFiles.UserId,
		ApiKey:     pageBehavior.PublicFiles.ApiKey,
		HappenTime: pageBehavior.PublicFiles.HappenTime,
		ActionType: pageBehavior.PublicFiles.ActionType,
		ActionID:   pageBehavior.ID,
		EventId:    eventId,
		ClassName:  pageBehavior.ClassName,
		InnterText: pageBehavior.InnterText,
	}
	CreateUserAction(userActionModel)
	return err
}

func CreatePageJsError(pageJsError model.PageJsError, eventId string) error {
	if err := global.GVA_DB.Create(&pageJsError).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:    pageJsError.PageUrl,
		UserId:     pageJsError.PublicFiles.UserId,
		ApiKey:     pageJsError.PublicFiles.ApiKey,
		HappenTime: pageJsError.PublicFiles.HappenTime,
		HappenDay:  pageJsError.PublicFiles.HappenDay,
		ActionType: pageJsError.PublicFiles.ActionType,
		ActionID:   pageJsError.ID,
		Message:    pageJsError.Message,
		Stack:      "pageJsError.Stack",
		EventId:    eventId,
	}
	CreateUserAction(userActionModel)
	return nil
}

func CreatePageView(pageView model.PageView, eventId string) error {
	if err := global.GVA_DB.Create(&pageView).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:    pageView.PageUrl,
		UserId:     pageView.PublicFiles.UserId,
		ApiKey:     pageView.PublicFiles.ApiKey,
		HappenTime: pageView.PublicFiles.HappenTime,
		HappenDay:  pageView.PublicFiles.HappenDay,
		ActionType: pageView.PublicFiles.ActionType,
		ActionID:   pageView.ID,
		EventId:    eventId,
	}
	err := CreateUserAction(userActionModel)
	return err
}

//func getTodayStartAndEndTime() (startTime string, endTime string) {
//	startTime = time.Now().Format("2006-01-02 00:00")
//	endTime = (time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Now().Location())).Format("2006-01-02 15:04:05")
//	return startTime, endTime
//}

// 获取资源错误
//func GetWebResourceErrorInfo() *response.PageResourcesResponse {
//	var resourcesList []response.ResourcesListResponse
//	var resourcesQuota response.ResourcesQuotaResponse
//	startTime, endTime := getTodayStartAndEndTime()
//
//	err := global.GVA_DB.Model(&model.PageResourceError{}).Select("source_url AS page_source_url, "+
//		"COUNT( source_url ) AS source_count, "+
//		"COUNT( DISTINCT user_id ) user_count, "+
//		"element_type, "+
//		"( SELECT COUNT( DISTINCT page_url ) AS page_url_count FROM page_resource_errors WHERE page_resource_errors.source_url = page_source_url ) AS page_url_count"+
//		"").Where("from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_source_url").Find(&resourcesList)
//
//	err = global.GVA_DB.Model(&model.PageResourceError{}).Select(" COUNT(*) as error_count,"+
//		"COUNT(page_url) as error_page, "+
//		"COUNT(DISTINCT user_id) as error_user").Where("from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&resourcesQuota)
//	fmt.Print(err, "err!")
//	return &response.PageResourcesResponse{
//		ResourcesList:  resourcesList,
//		ResourcesQuota: resourcesQuota,
//	}
//}
