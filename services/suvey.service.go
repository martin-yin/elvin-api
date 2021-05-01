package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func GetSurveyStatisticsData(startTime string, endTime string, monitorId string) (surveyStatisticsResponse response.SurveyStatisticsResponse, err error) {
	fmt.Println(startTime, "startTime-----")
	fmt.Println(endTime, "endTime-----")

	//sqlWhere := `from_unixtime(happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select(" round( AVG( load_page ), 2 ) AS load_page").Where(`from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`, startTime, endTime, monitorId).Find(&surveyStatisticsResponse.LoadPage).Error
	err = global.GVA_DB.Model(&model.PageResourceError{}).Select("COUNT( DISTINCT id ) as resources").Where(`from_unixtime(page_resources_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`, startTime, endTime, monitorId).Find(&surveyStatisticsResponse.Resources).Error
	err = global.GVA_DB.Model(&model.PageJsError{}).Select("COUNT( DISTINCT id ) as js_error").Where(`from_unixtime(page_js_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`, startTime, endTime, monitorId).Find(&surveyStatisticsResponse.JsError).Error
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("COUNT( DISTINCT id ) as http_error").Where(`from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?` + " AND page_https.status >= 400", startTime, endTime, monitorId).Find(&surveyStatisticsResponse.HttpError).Error
	return
}

func GetSurveyPUvData(startTime string, endTime string, monitorId string) (surveyPUvData []response.SurveyPUvData, err error) {
	sqlWhere := "from_unixtime(page_views.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? "
	err = global.GVA_DB.Model(&model.PageView{}).Select("FROM_UNIXTIME( happen_time / 1000, '%H:%i') AS time_key, COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(sqlWhere, startTime, endTime, monitorId).Group("time_key").Scan(&surveyPUvData).Error
	return
}

