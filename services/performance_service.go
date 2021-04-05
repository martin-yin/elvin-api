package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetStackPerformance(startTime string, endTime string) (stackData response.StackResponse, err error) {
	err = global.GVA_DB.Model(&model.LoadpageInfo{}).Select("round(AVG(redirect),2) as redirect, "+
		"round(AVG(appcache),2) as appcache, "+
		"round(AVG(lookup_domain),2) as lookup_domain, "+
		"round(AVG(tcp),2) as tcp, "+
		"round(AVG(ssl_t),2) as ssl_t, "+
		"round(AVG(request),2) as request, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_event),2) as load_event, "+
		"round(AVG(dom_parse),2) as dom_parse ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&stackData).Error
	return
}

func GetQuotaData(startTime string, endTime string) (quotaData response.QuotaResponse, err error) {
	err = global.GVA_DB.Model(&model.LoadpageInfo{}).Select("round(AVG(dom_parse),2) as dom_parse, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_page),2) as load_page, "+
		"Count(*) as Pv ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quotaData).Error
	err = global.GVA_DB.Model(&model.LoadpageInfo{}).Select("CONCAT(round((SELECT COUNT(*) as pv FROM loadpage_infos WHERE loadpage_infos.load_page < 2000) / Count( * ) * 100, 2), '%')  AS Score").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quotaData.Fast).Error
	return
}

func GetStageTimeList(startTime string, endTime string) (stageTimeList []response.StageTimeResponse, err error) {
	err = global.GVA_DB.Model(&model.LoadpageInfo{}).Select("DATE_FORMAT(happen_time, \"%H:%i\") AS time_key, "+
		"round( AVG( redirect ), 2 ) AS redirect,"+
		"round( AVG( appcache ), 2 ) AS appcache,"+
		"round( AVG( lookup_domain ), 2 ) AS lookup_domain,"+
		"round( AVG( tcp ), 2 ) AS tcp,"+
		"round( AVG( ssl_t ), 2 ) AS ssl_t,"+
		"round( AVG( request ), 2 ) AS request,"+
		"round( AVG( ttfb ), 2 ) AS ttfb,"+
		"round( AVG( load_event ), 2 ) AS load_event,"+
		"round( AVG( load_page ), 2 ) AS load_page,"+
		"COUNT(*) as Pv ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("time_key").Scan(&stageTimeList).Error
	return
}

func GetLoadInfoPageList(startTime string, endTime string) (loadpageInfoList []response.LoadpageInfoListResponse, err error) {
	err = global.GVA_DB.Model(&model.LoadpageInfo{}).Select("ID, page_url, "+
		"request, "+
		"dom_parse, "+
		"ttfb, "+
		"load_page, "+
		"load_event, "+
		"load_type, "+
		"load_page, "+
		"COUNT(*) as pv ").Where("loadpage_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_url").Scan(&loadpageInfoList).Error
	return
}
type UserInfo struct {
	PageUrl      string  `json:"page_url"`
	UserId       string  `json:"user_id"`
	ApiKey       string  `json:"api_key"`
	UploadType   string  `json:"upload_type"`
	HappenTime   string  `json:"happen_time"`
	Redirect     float64 `json:"redirect"`
	Appcache     float64 `json:"appcache"`
	LookupDomain float64 `json:"lookup_domain"`
	Tcp          float64 `json:"tcp"`
	SslT         float64 `json:"ssl_t"`
	Request      float64 `json:"request"`
	DomParse     float64 `json:"dom_parse"`
	Ttfb         float64 `json:"ttfb"`
	LoadPage     float64 `json:"load_page"`
	LoadEvent    float64 `json:"load_event"`
	LoadType     string  `json:"load_type"`
}

func GetPerformance(id string) (userInfo UserInfo, err error) {
	err = global.GVA_DB.Model(&model.LoadpageInfo{}).Where("id = ?", id).Scan(&userInfo).Error
	return
}
