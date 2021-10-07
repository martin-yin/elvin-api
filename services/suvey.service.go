package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/response"
	"fmt"
)

func GetSurveyStatisticsData(startTime string, endTime string, monitorId string) (surveyStatisticsResponse response.SurveyStatisticsResponse, err error) {
	fmt.Println(startTime, "startTime-----")
	fmt.Println(endTime, "endTime-----")

	//sqlWhere := `from_unixtime(happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`
	err = global.GORMDB.Model(&model.PagePerformance{}).Select(" round( AVG( load_page ), 2 ) AS load_page").Where(`from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`, startTime, endTime, monitorId).Find(&surveyStatisticsResponse.LoadPage).Error
	err = global.GORMDB.Model(&model.PageResourceError{}).Select("COUNT( DISTINCT id ) as resources").Where(`from_unixtime(page_resource_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`, startTime, endTime, monitorId).Find(&surveyStatisticsResponse.Resources).Error
	err = global.GORMDB.Model(&model.Issue{}).Select("COUNT( DISTINCT id ) as js_error").Where(`from_unixtime(issues.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`, startTime, endTime, monitorId).Find(&surveyStatisticsResponse.JsError).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select("COUNT( DISTINCT id ) as http_error").Where(`from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`+" AND page_https.status >= 400", startTime, endTime, monitorId).Find(&surveyStatisticsResponse.HttpError).Error
	return
}

func GetSurveyPUvData(startTime string, endTime string, monitorId string) (surveyPUvData []response.SurveyPUvData, err error) {
	sqlWhere := "from_unixtime(page_views.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? "
	err = global.GORMDB.Model(&model.PageView{}).Select("FROM_UNIXTIME( happen_time / 1000, '%H:00') AS time_key, COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(sqlWhere, startTime, endTime, monitorId).Group("time_key").Find(&surveyPUvData).Error
	return
}

func GetSurveyJsErrorData(startTime string, endTime string, monitorId string) (surveyJsErrorData []response.JsErrorData, err error) {
	sqlWhere := "from_unixtime(page_js_errors.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? "
	err = global.GORMDB.Model(&model.PageView{}).Select("FROM_UNIXTIME( happen_time / 1000, '%H:%i') AS time_key, COUNT( DISTINCT user_id ) as user, COUNT( DISTINCT id ) as js_error").Where(sqlWhere, startTime, endTime, monitorId).Group("time_key").Scan(&surveyJsErrorData).Error
	return
}
