package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
	"reflect"
	"time"
)

func CreatePagePerformance(pagePerformance model.PagePerformance, eventId string) error {
	if err := global.GVA_DB.Create(&pagePerformance).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:    pagePerformance.PageUrl,
		UserId:     pagePerformance.UserId,
		ApiKey:     pagePerformance.ApiKey,
		HappenTime: pagePerformance.HappenTime,
		HappenDay:  pagePerformance.HappenDay,
		ActionType: pagePerformance.ActionType,
		LoadType:   pagePerformance.LoadType,
		EventId:    eventId,
		ActionID:   pagePerformance.ID,
	}

	userModel := model.User{
		UserId:         pagePerformance.UserId,
		ApiKey:         pagePerformance.ApiKey,
		HappenTime:     pagePerformance.HappenTime,
		HappenDay:      pagePerformance.HappenDay,
		EventId:        eventId,
		IP:             pagePerformance.IP,
		Device:         pagePerformance.Device,
		DeviceType:     pagePerformance.DeviceType,
		Os:             pagePerformance.Os,
		OsVersion:      pagePerformance.OsVersion,
		Browser:        pagePerformance.Browser,
		BrowserVersion: pagePerformance.BrowserVersion,
		UA:             pagePerformance.UA,
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
		fmt.Print(err, "err !!!!!!!!!!!!")
		return err
	}
	return nil
}

// 存 http请求信息。
func CreatePageHttpModel(pageHttp model.PageHttp, eventId string) error {
	var pageHttpStatistical model.PageHttpStatistical
	err := global.GVA_DB.Model(&model.PageHttpStatistical{}).Where("http_url = ?", pageHttp.HttpUrl).Find(&pageHttpStatistical).Error
	if !reflect.DeepEqual(pageHttpStatistical, model.PageHttpStatistical{}) {
		pageHttpStatistical.Total++
		if pageHttp.Status > 200 {
			pageHttpStatistical.FailTotal++
		} else {
			pageHttpStatistical.SuccessTotal++
		}
	} else {
		pageHttpStatistical.PageUrl = pageHttp.PageUrl
		pageHttpStatistical.HttpUrl = pageHttp.HttpUrl
		pageHttpStatistical.Total++
		if pageHttp.Status > 200 {
			pageHttpStatistical.FailTotal++
		} else {
			pageHttpStatistical.SuccessTotal++
		}
	}
	err = global.GVA_DB.Save(&pageHttpStatistical).Error
	err = global.GVA_DB.Create(&pageHttp).Error

	userActionModel := model.UserAction{
		PageUrl:    pageHttp.PageUrl,
		UserId:     pageHttp.UserId,
		ApiKey:     pageHttp.ApiKey,
		EventId:    eventId,
		HappenTime: pageHttp.HappenTime,
		HappenDay:  pageHttp.HappenDay,
		ActionType: pageHttp.ActionType,
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
		UserId:      resourceErrorInfo.UserId,
		ApiKey:      resourceErrorInfo.ApiKey,
		HappenTime:  resourceErrorInfo.HappenTime,
		HappenDay:   resourceErrorInfo.HappenDay,
		ActionType:  resourceErrorInfo.ActionType,
		ActionID:    resourceErrorInfo.ID,
		EventId:     eventId,
		SourceUrl:   resourceErrorInfo.SourceUrl,
		ElementType: resourceErrorInfo.ElementType,
	}
	CreateUserAction(userActionModel)
	return nil
}

func CreatePageBehavior(pageBehavior model.PageBehavior, eventId string) error {
	if err := global.GVA_DB.Create(&pageBehavior).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:    pageBehavior.PageUrl,
		UserId:     pageBehavior.UserId,
		ApiKey:     pageBehavior.ApiKey,
		HappenTime: pageBehavior.HappenTime,
		ActionType: pageBehavior.ActionType,
		ActionID:   pageBehavior.ID,
		EventId:    eventId,
		ClassName:  pageBehavior.ClassName,
		InnterText: pageBehavior.InnterText,
	}
	CreateUserAction(userActionModel)
	return nil
}

func CreatePageJsError(pageJsError model.PageJsError, eventId string) error {
	if err := global.GVA_DB.Create(&pageJsError).Error; err != nil {
		return err
	}
	userActionModel := model.UserAction{
		PageUrl:    pageJsError.PageUrl,
		UserId:     pageJsError.UserId,
		ApiKey:     pageJsError.ApiKey,
		HappenTime: pageJsError.HappenTime,
		HappenDay:  pageJsError.HappenDay,
		ActionType: pageJsError.ActionType,
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
		UserId:     pageView.UserId,
		ApiKey:     pageView.ApiKey,
		HappenTime: pageView.HappenTime,
		HappenDay:  pageView.HappenDay,
		ActionType: pageView.ActionType,
		ActionID:   pageView.ID,
		EventId:    eventId,
	}
	err := CreateUserAction(userActionModel)
	fmt.Print(err, "!!!!!!!!!!!!!!!!!!")
	return nil
}

func getTodayStartAndEndTime() (startTime string, endTime string) {
	startTime = time.Now().Format("2006-01-02 00:00")
	endTime = (time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Now().Location())).Format("2006-01-02 15:04:05")
	return startTime, endTime
}

// 获取资源错误
func GetWebResourceErrorInfo() *response.PageResourcesResponse {
	var resourcesList []response.ResourcesListResponse
	var resourcesQuota response.ResourcesQuotaResponse
	startTime, endTime := getTodayStartAndEndTime()

	err := global.GVA_DB.Model(&model.PageResourceError{}).Select("source_url AS page_source_url, "+
		"COUNT( source_url ) AS source_count, "+
		"COUNT( DISTINCT user_id ) user_count, "+
		"element_type, "+
		"( SELECT COUNT( DISTINCT page_url ) AS page_url_count FROM page_resource_errors WHERE page_resource_errors.source_url = page_source_url ) AS page_url_count"+
		"").Where("from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_source_url").Find(&resourcesList)

	err = global.GVA_DB.Model(&model.PageResourceError{}).Select(" COUNT(*) as error_count,"+
		"COUNT(page_url) as error_page, "+
		"COUNT(DISTINCT user_id) as error_user").Where("from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&resourcesQuota)
	fmt.Print(err, "err!")
	return &response.PageResourcesResponse{
		ResourcesList:  resourcesList,
		ResourcesQuota: resourcesQuota,
	}
}
