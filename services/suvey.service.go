package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetSurveyStatisticsData(startTime string, endTime string, monitorId string) (surveyStatisticsResponse response.SurveyStatisticsResponse, err error) {
	sqlWhere := `from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?`
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT ip ) as ip").Where(sqlWhere, startTime, endTime, monitorId).Find(&surveyStatisticsResponse).Error
	err = global.GVA_DB.Model(&model.PageView{}).Select("COUNT(*) as pv").Where("from_unixtime(page_views.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')  and monitor_id = ?", startTime, endTime, monitorId).Find(&surveyStatisticsResponse).Error
	return
}

func GetSurveyPerformance(startTime string, endTime string, monitorId string) (surveyPerformancesResponse []response.SurveyPerformancesResponse, err error) {
	sqlWhere := "from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? "
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("FROM_UNIXTIME( happen_time / 1000, '%H:%i') AS time_key, round( AVG( load_page ), 2 ) AS load_page").Where(sqlWhere, startTime, endTime, monitorId).Group("time_key").Find(&surveyPerformancesResponse).Error
	return
}

func GetHttp(startTime string, endTime string, monitorId string) {

}
