package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func SetWebLoadPageInfo(weLoadPageInfo model.WebLoadpageInfo) {

	err := global.GVA_DB.Create(&weLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func WebHttpInfoModel(weLoadPageInfo []*model.WebHttpInfo) {
	err := global.GVA_DB.Create(&weLoadPageInfo).Error
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

// 获取 web load 得信息
func GetWebLoadPageInfo() *response.WebLoadPageInfoResponse {
	var stacking response.StackResponse
	var quota response.QuotaResponse
	var loadpageInfoList []response.LoadpageInfoListResponse

	// 这里不赋值给一个变量是因为有bug！！！
	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("ID, page_url, request, dom_parse, ttfb, load_page, load_event, load_type, load_page, COUNT(page_url) as pv").Group("page_url").Scan(&loadpageInfoList)
	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("round(AVG(redirect),2) as redirect,  round(AVG(appcache),2) as appcache, round(AVG(lookup_domain),2) as lookup_domain,  round(AVG(tcp),2) as tcp, round(AVG(ssl_t),2) as ssl_t,  round(AVG(request),2) as request, round(AVG(ttfb),2) as ttfb, round(AVG(load_event),2) as load_event, round(AVG(dom_parse),2) as dom_parse").Find(&stacking)
	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("round(AVG(dom_parse),2) as dom_parse, round(AVG(ttfb),2) as ttfb, round(AVG(load_page),2) as load_page, Count(id) as Pv").Scan(&quota)
	global.GVA_DB.Model(&model.WebLoadpageInfo{}).Select("CONCAT(round((SELECT COUNT(id) as pv  FROM web_loadpage_infos WHERE web_loadpage_infos.load_page < 2000) / Count( id ) * 100, 2), '%')  AS Score ").Scan(&quota.Fast)

	return &response.WebLoadPageInfoResponse{
		QuotaResponse: quota,
		StackResponse: stacking,
		LoadpageInfoListResponse:  loadpageInfoList,
	}
}

func GetWebHttpInfo() *response.WebHttpInfoListResponse {
	var httpInfoList []response.HttpInfoListResponse
	var httpQuota response.HttpQuotaResponse

	err := global.GVA_DB.Model(&model.WebHttpInfo{}).Select("http_url AS http_url, "+
		"page_url, " +
		"COUNT( http_url ) AS request_total, " +
		"round(AVG( load_time ), 2) AS load_time, " +
		"CONCAT(round(( SELECT COUNT( http_url ) FROM web_http_infos WHERE http_url = request_url AND web_http_infos.`status` != 0 AND web_http_infos.`status` BETWEEN 200 AND 305 ) / COUNT( http_url )  * 100, 2), '%')  AS success_rate ").Where("web_http_infos.`status` != 0 ").Group("request_url").Find(&httpInfoList)
	fmt.Print(err, "err!")
	global.GVA_DB.Model(&model.WebHttpInfo{}).Select("" +
		"COUNT( http_url ) AS request_total, " +
		"round(AVG( load_time )) AS load_time, " +
		"CONCAT(round((SELECT COUNT( http_url ) FROM web_http_infos WHERE web_http_infos.`status` != 0 AND web_http_infos.`status` BETWEEN 200 AND 305 ) / COUNT( http_url ) * 100,2 ), '%') AS success_rate," +
		"(SELECT COUNT(DISTINCT user_id) as user FROM web_http_infos  WHERE web_http_infos.`status` > 305 ) as error_user").Where("web_http_infos.`status` != 0").Find(&httpQuota)

	return &response.WebHttpInfoListResponse{
		HttpInfoListResponse: httpInfoList,
		HttpQuotaResponse: httpQuota,
	}

	//请求次数
	// 成功率
	// 请求耗时
	// 失败影响用户

	// 请求成功列表
}
