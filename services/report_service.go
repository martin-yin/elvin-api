package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
	"reflect"
	"time"
)

func CreateLoadPageInfo(weLoadPageInfo model.LoadpageInfo) error {
	if err := global.GVA_DB.Create(&weLoadPageInfo).Error; err != nil {
		return err
	}
	return nil
}

// 存 http请求信息。
func CreateHttpInfoModel(webLoadPageInfo []model.HttpInfo) error {
	var webhttpInfoInfoStatistical model.HttpInfoStatistical
	err := global.GVA_DB.Model(&model.HttpInfoStatistical{}).Where("http_url = ?", webLoadPageInfo[1].HttpUrl).Find(&webhttpInfoInfoStatistical).Error
	if !reflect.DeepEqual(webhttpInfoInfoStatistical, model.HttpInfoStatistical{}) {
		webhttpInfoInfoStatistical.Total++
		if webLoadPageInfo[1].Status > 200 {
			webhttpInfoInfoStatistical.FailTotal++
		} else {
			webhttpInfoInfoStatistical.SuccessTotal++
		}
	} else {
		webhttpInfoInfoStatistical.PageUrl = webLoadPageInfo[1].PageUrl
		webhttpInfoInfoStatistical.HttpUrl = webLoadPageInfo[1].HttpUrl
		webhttpInfoInfoStatistical.Total++
		if webLoadPageInfo[1].Status > 200 {
			webhttpInfoInfoStatistical.FailTotal++
		} else {
			webhttpInfoInfoStatistical.SuccessTotal++
		}
	}
	err = global.GVA_DB.Save(&webhttpInfoInfoStatistical).Error
	err = global.GVA_DB.Create(&webLoadPageInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateResourcesError(webResourceErrorInfo model.ResourceErrorInfo) error {
	if err := global.GVA_DB.Create(&webResourceErrorInfo).Error; err != nil {
		return err
	}
	return nil
}

func CreateBehaviorInfo(webBehaviorInfo model.BehaviorInfo) error {
	if err := global.GVA_DB.Create(&webBehaviorInfo).Error; err != nil {
		return err
	}
	return nil
}

func getTodayStartAndEndTime() (startTime string, endTime string) {
	startTime = time.Now().Format("2006-01-02 00:00")
	endTime = (time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Now().Location())).Format("2006-01-02 15:04:05")
	return startTime, endTime
}

// 获取 web load 得信息
func GetWebLoadPageInfo() *response.WebLoadPageInfoResponse {
	var stacking response.StackResponse
	var quota response.QuotaResponse
	var loadpageInfoList []response.LoadpageInfoListResponse
	var stageTime []response.StageTimeResponse
	startTime, endTime := getTodayStartAndEndTime()

	// 这里不赋值给一个变量是因为有bug！！！
	global.GVA_DB.Model(&model.LoadpageInfo{}).Select("ID, page_url, "+
		"request, "+
		"dom_parse, "+
		"ttfb, "+
		"load_page, "+
		"load_event, "+
		"load_type, "+
		"load_page, "+
		"COUNT(*) as pv ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_url").Scan(&loadpageInfoList)

	global.GVA_DB.Model(&model.LoadpageInfo{}).Select("round(AVG(redirect),2) as redirect, "+
		"round(AVG(appcache),2) as appcache, "+
		"round(AVG(lookup_domain),2) as lookup_domain, "+
		"round(AVG(tcp),2) as tcp, "+
		"round(AVG(ssl_t),2) as ssl_t, "+
		"round(AVG(request),2) as request, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_event),2) as load_event, "+
		"round(AVG(dom_parse),2) as dom_parse ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&stacking)

	global.GVA_DB.Model(&model.LoadpageInfo{}).Select("round(AVG(dom_parse),2) as dom_parse, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_page),2) as load_page, "+
		"Count(*) as Pv ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quota)

	global.GVA_DB.Model(&model.LoadpageInfo{}).Select("CONCAT(round((SELECT COUNT(*) as pv FROM loadpage_infos WHERE loadpage_infos.load_page < 2000) / Count( * ) * 100, 2), '%')  AS Score").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quota.Fast)

	global.GVA_DB.Model(&model.LoadpageInfo{}).Select("DATE_FORMAT(happen_time, \"%H:%i\") AS time_key, "+
		"round( AVG( redirect ), 2 ) AS redirect,"+
		"round( AVG( appcache ), 2 ) AS appcache,"+
		"round( AVG( lookup_domain ), 2 ) AS lookup_domain,"+
		"round( AVG( tcp ), 2 ) AS tcp,"+
		"round( AVG( ssl_t ), 2 ) AS ssl_t,"+
		"round( AVG( request ), 2 ) AS request,"+
		"round( AVG( ttfb ), 2 ) AS ttfb,"+
		"round( AVG( load_event ), 2 ) AS load_event,"+
		"round( AVG( load_page ), 2 ) AS load_page,"+
		"COUNT(*) as Pv ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("time_key").Scan(&stageTime)

	return &response.WebLoadPageInfoResponse{
		QuotaResponse:            quota,
		StackResponse:            stacking,
		LoadpageInfoListResponse: loadpageInfoList,
		StageTimeResponse:        stageTime,
	}
}

// 获取http请求
func GetWebHttpInfo() *response.WebHttpInfoResponse {
	var httpInfoList []response.HttpInfoListResponse
	var httpQuota response.HttpQuotaResponse
	startTime, endTime := getTodayStartAndEndTime()

	err := global.GVA_DB.Model(&model.HttpInfo{}).Select("http_infos.http_url, "+
		"http_infos.page_url, "+
		"round( AVG( load_time ), 2 ) AS load_time, "+
		"total, fail_total, success_total").Where("http_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Joins("LEFT JOIN http_info_statisticals statisticals ON statisticals.http_url = http_infos.http_url").Where("http_infos.`status` != 0 ").Group("http_url").Find(&httpInfoList)

	err = global.GVA_DB.Model(&model.HttpInfo{}).Select(""+
		"COUNT( * ) AS total, "+
		"round(AVG( load_time )) AS load_time, "+
		"(SELECT COUNT(DISTINCT user_id) as user FROM http_infos  WHERE http_infos.`status` > 305 ) as error_user, "+
		"( SELECT COUNT(  * ) FROM http_infos WHERE http_infos.`status` BETWEEN 200 AND 305 ) AS success_total").Where("http_infos.`status` != 0 And http_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&httpQuota)
	fmt.Print(err, "err \n")
	return &response.WebHttpInfoResponse{
		HttpInfoListResponse: httpInfoList,
		HttpQuotaResponse:    httpQuota,
	}
}

// 获取资源错误
func GetWebResourceErrorInfo() *response.WebResourcesInfoResponse {
	var resourcesInfoList []response.ResourcesInfoListResponse
	var resourcesQuota response.ResourcesQuota
	startTime, endTime := getTodayStartAndEndTime()

	err := global.GVA_DB.Model(&model.ResourceErrorInfo{}).Select("source_url AS page_source_url, "+
		"COUNT( source_url ) AS source_count, "+
		"COUNT( DISTINCT user_id ) user_count, "+
		"element_type, "+
		"( SELECT COUNT( DISTINCT page_url ) AS page_url_count FROM resource_error_infos WHERE resource_error_infos.source_url = page_source_url ) AS page_url_count"+
		"").Where("resource_error_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_source_url").Find(&resourcesInfoList)

	err = global.GVA_DB.Model(&model.ResourceErrorInfo{}).Select(" COUNT(*) as error_count,"+
		"COUNT(page_url) as error_page, "+
		"COUNT(DISTINCT user_id) as error_user").Where("resource_error_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&resourcesQuota)
	fmt.Print(err, "err!")
	return &response.WebResourcesInfoResponse{
		ResourcesInfoList: resourcesInfoList,
		ResourcesQuota:    resourcesQuota,
	}
}
