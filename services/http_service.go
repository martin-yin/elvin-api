package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetHttpInfoList(startTime string, endTime string) (httpInfoList []response.HttpInfoListResponse, err error) {
	err = global.GVA_DB.Model(&model.HttpInfo{}).Select("http_infos.http_url, "+
		"http_infos.page_url, "+
		"round( AVG( load_time ), 2 ) AS load_time, "+
		"total, fail_total, success_total").Where("http_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Joins("LEFT JOIN http_info_statisticals statisticals ON statisticals.http_url = http_infos.http_url").Where("http_infos.`status` != 0 ").Group("http_url").Find(&httpInfoList).Error
	return
}

func GetHttpQuota(startTime string, endTime string) (httpQuota response.HttpQuotaResponse, err error) {
	err = global.GVA_DB.Model(&model.HttpInfo{}).Select(""+
		"COUNT( * ) AS total, "+
		"round(AVG( load_time )) AS load_time, "+
		"(SELECT COUNT(DISTINCT user_id) as user FROM http_infos  WHERE http_infos.`status` > 305 ) as error_user, "+
		"( SELECT COUNT(  * ) FROM http_infos WHERE http_infos.`status` BETWEEN 200 AND 305 ) AS success_total").Where("http_infos.`status` != 0 And http_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&httpQuota).Error
	return
}
