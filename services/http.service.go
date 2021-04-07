package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetHttpInfoList(startTime string, endTime string) (httpInfoList []response.HttpListResponse, err error) {
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("page_https.http_url, "+
		"page_https.page_url, "+
		"round( AVG( load_time ), 2 ) AS load_time, "+
		"total, fail_total, success_total").Where("from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Joins("LEFT JOIN page_http_statisticals statisticals ON statisticals.http_url = page_https.http_url").Where("page_https.`status` != 0 ").Group("http_url").Find(&httpInfoList).Error
	return
}

func GetHttpQuota(startTime string, endTime string) (httpQuota response.HttpQuotaResponse, err error) {
	err = global.GVA_DB.Model(&model.PageHttp{}).Select(""+
		"COUNT( * ) AS total, "+
		"round(AVG( load_time )) AS load_time, "+
		"(SELECT COUNT(DISTINCT user_id) as user FROM page_https  WHERE page_https.`status` > 305 ) as error_user, "+
		"( SELECT COUNT(  * ) FROM page_https WHERE page_https.`status` BETWEEN 200 AND 305 ) AS success_total").Where("page_https.`status` != 0 And from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&httpQuota).Error
	return
}
