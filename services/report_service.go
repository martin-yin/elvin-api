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

func CreateJsErrorInfo(jsErrorInfO model.JsErrorInfo) error {
	if err := global.GVA_DB.Create(&jsErrorInfO).Error; err != nil {
		return err
	}
	return nil
}

func getTodayStartAndEndTime() (startTime string, endTime string) {
	startTime = time.Now().Format("2006-01-02 00:00")
	endTime = (time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Now().Location())).Format("2006-01-02 15:04:05")
	return startTime, endTime
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
