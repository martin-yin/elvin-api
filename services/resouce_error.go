package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetResourcesInfoList(startTime string, endTime string) (resourcesInfoList []response.ResourcesInfoListResponse, err error) {
	err = global.GVA_DB.Model(&model.ResourceErrorInfo{}).Select("source_url AS page_source_url, "+
		"COUNT( source_url ) AS source_count, "+
		"COUNT( DISTINCT user_id ) user_count, "+
		"element_type, "+
		"( SELECT COUNT( DISTINCT page_url ) AS page_url_count FROM resource_error_infos WHERE resource_error_infos.source_url = page_source_url ) AS page_url_count"+
		"").Where("resource_error_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_source_url").Find(&resourcesInfoList).Error
	return
}

func GetResourcesQuota(startTime string, endTime string) (resourcesQuota response.ResourcesQuota, err error) {
	err = global.GVA_DB.Model(&model.ResourceErrorInfo{}).Select(" COUNT(*) as error_count,"+
		"COUNT(page_url) as error_page, "+
		"COUNT(DISTINCT user_id) as error_user").Where("resource_error_infos.happen_time between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&resourcesQuota).Error
	return
}
