package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetResourcesInfoList(monitorId string, startTime string, endTime string) (resourcesList []response.ResourcesListResponse, err error) {
	err = global.GVA_DB.Model(&model.PageResourceError{}).Select("source_url AS page_source_url, "+
		"COUNT( source_url ) AS source_count, "+
		"COUNT( DISTINCT user_id ) user_count, "+
		"element_type, "+
		"( SELECT COUNT( DISTINCT page_url ) AS page_url_count FROM page_resource_errors WHERE page_resource_errors.source_url = page_source_url ) AS page_url_count"+
		"").Where("from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ?", startTime, endTime, monitorId).Group("page_source_url").Find(&resourcesList).Error
	return
}

func GetResourcesQuota(monitorId string, startTime string, endTime string) (resourcesQuota response.ResourcesQuotaResponse, err error) {
	err = global.GVA_DB.Model(&model.PageResourceError{}).Select(" COUNT(*) as error_count,"+
		"COUNT(DISTINCT page_url) as error_page, "+
		"COUNT(DISTINCT user_id) as error_user").Where("from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ?", startTime, endTime, monitorId).Find(&resourcesQuota).Error
	return
}
