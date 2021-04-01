package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
	"reflect"
	"time"
)

func SetWebLoadPageInfo(weLoadPageInfo model.WebLoadpageInfo) {
	err := global.GVA_DB.Create(&weLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

// 存 http请求信息。
func WebHttpInfoModel(webLoadPageInfo []*model.WebHttpInfo) {
	var webhttpInfoInfoStatistical model.WebHttpInfoStatistical
	global.GVA_DB.Model(&model.WebHttpInfoStatistical{}).Where("http_url = ?", webLoadPageInfo[1].HttpUrl).Find(&webhttpInfoInfoStatistical)
	// HttpInfoStatistical
	//var httpErrorNumber int = 0;

	if !reflect.DeepEqual(webhttpInfoInfoStatistical, model.WebHttpInfoStatistical{}) {
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
	global.GVA_DB.Save(&webhttpInfoInfoStatistical)
	err := global.GVA_DB.Create(&webLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func SetWebResourcesError(webResourceErrorInfo model.WebResourceErrorInfo) {
	err := global.GVA_DB.Create(&webResourceErrorInfo).Error
	if err != nil {
		fmt.Print(err, "err !!!!!!!! \n")
	}
}

func SetBehaviorInfo(webBehaviorInfo model.WebBehaviorInfo) {
	err := global.GVA_DB.Create(&webBehaviorInfo).Error
	if err != nil {
		fmt.Print(err, "err !!!!!!!! \n")
	}
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

	//fmt.Print(, "getTody \n")
	// 这里不赋值给一个变量是因为有bug！！！
	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("ID, page_url, "+
		"request, "+
		"dom_parse, "+
		"ttfb, "+
		"load_page, "+
		"load_event, "+
		"load_type, "+
		"load_page, "+
		"COUNT(page_url) as pv ").Where("web_loadpage_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_url").Scan(&loadpageInfoList)

	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("round(AVG(redirect),2) as redirect, "+
		"round(AVG(appcache),2) as appcache, "+
		"round(AVG(lookup_domain),2) as lookup_domain, "+
		"round(AVG(tcp),2) as tcp, "+
		"round(AVG(ssl_t),2) as ssl_t, "+
		"round(AVG(request),2) as request, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_event),2) as load_event, "+
		"round(AVG(dom_parse),2) as dom_parse ").Where("web_loadpage_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&stacking)

	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("round(AVG(dom_parse),2) as dom_parse, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_page),2) as load_page, "+
		"Count(id) as Pv ").Where("web_loadpage_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quota)

	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("CONCAT(round((SELECT COUNT(id) as pv FROM web_loadpage_infos WHERE web_loadpage_infos.load_page < 2000) / Count( id ) * 100, 2), '%')  AS Score").Where("web_loadpage_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quota.Fast)

	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("FROM_UNIXTIME ( happen_time / 1000, \"%H:%i\" ) AS time_key, "+
		"round( AVG( redirect ), 2 ) AS redirect,"+
		"round( AVG( appcache ), 2 ) AS appcache,"+
		"round( AVG( lookup_domain ), 2 ) AS lookup_domain,"+
		"round( AVG( tcp ), 2 ) AS tcp,"+
		"round( AVG( ssl_t ), 2 ) AS ssl_t,"+
		"round( AVG( request ), 2 ) AS request,"+
		"round( AVG( ttfb ), 2 ) AS ttfb,"+
		"round( AVG( load_event ), 2 ) AS load_event,"+
		"round( AVG( load_page ), 2 ) AS load_page,"+
		"COUNT(*) as Pv ").Where("web_loadpage_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("time_key").Scan(&stageTime)

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

	err := global.GVA_DB.Model(&model.WebHttpInfo{}).Select("http_url AS request_url, " +
		"page_url, " +
		"COUNT( http_url ) AS request_total, " +
		"round(AVG( load_time ), 2) AS load_time, " +
		"CONCAT(round(( SELECT COUNT( http_url ) FROM web_http_infos WHERE http_url = request_url AND web_http_infos.`status` != 0 AND web_http_infos.`status` BETWEEN 200 AND 305 ) / COUNT( http_url )  * 100, 2), '%')  AS success_rate").Where("web_http_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Where("web_http_infos.`status` != 0 ").Group("request_url").Find(&httpInfoList)
	fmt.Print(err, "err \n")
	err = global.GVA_DB.Model(&model.WebHttpInfo{}).Select("" +
		"COUNT( http_url ) AS request_total, " +
		"round(AVG( load_time )) AS load_time, " +
		"(SELECT COUNT(DISTINCT user_id) as user FROM web_http_infos  WHERE web_http_infos.`status` > 305 ) as error_user").Where("web_http_infos.`status` != 0 And web_http_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&httpQuota)
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

	err := global.GVA_DB.Model(&model.WebResourceErrorInfo{}).Select("source_url AS page_source_url, " +
		"COUNT( source_url ) AS source_count, " +
		"COUNT( DISTINCT user_id ) user_count, " +
		"element_type, " +
		"( SELECT COUNT( DISTINCT page_url ) AS page_url_count FROM web_resource_error_infos WHERE web_resource_error_infos.source_url = page_source_url ) AS page_url_count" +
		"").Where("web_resource_error_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_source_url").Find(&resourcesInfoList)

	err = global.GVA_DB.Model(&model.WebResourceErrorInfo{}).Select(" COUNT(*) as error_count," +
		"COUNT(page_url) as error_page, " +
		"COUNT(DISTINCT user_id) as error_user").Where("web_resource_error_infos.created_at between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&resourcesQuota)
	fmt.Print(err, "err!")
	return &response.WebResourcesInfoResponse{
		ResourcesInfoList: resourcesInfoList,
		ResourcesQuota:    resourcesQuota,
	}

}
