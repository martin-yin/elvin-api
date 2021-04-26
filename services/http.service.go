package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
	"strconv"
)

func GetHttpInfoList(monitorId string, startTime string, endTime string) (httpInfoList []response.HttpListResponse, err error) {
	sqlWhere := `from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? and status = 200`
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("http_url, Count(DISTINCT http_url) as total, Count(DISTINCT user_id) as user_total, status, round( AVG( load_time ), 2 ) AS load_time").Where(sqlWhere, startTimes, endTimes, monitorId).Group("http_url").Find(&httpInfoList).Debug().Error
	return
}

func GetHttpQuota(monitorId string, startTime string, endTime string) (httpQuota response.HttpQuotaResponse, err error) {
	sqlWhere := `from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ?  `
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("round(AVG( load_time )) AS load_time, status, COUNT(*) AS total").Where(sqlWhere, startTimes, endTimes, monitorId).Find(&httpQuota).Error
	// 查询错误影响的用户
	err = global.GVA_DB.Model(&model.PageHttp{}).Select(" Count(DISTINCT user_id) as error_user, status ").Where(sqlWhere+" and status > 400", startTimes, endTimes, monitorId).Find(&httpQuota).Error
	// 查询成功的条数
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("COUNT(*) AS success_total, status").Where(sqlWhere+" and status < 400", startTimes, endTimes, monitorId).Scan(&httpQuota).Error
	return
}

func GetHttpStageTimeByTimeGrain(monitorId string, timeGrain string, startTime string, endTime string) (httpStageTime []response.HttpStageTimeResponse, err error) {
	query := ""
	sqlWhere := "from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? and status = 200"
	if timeGrain == "minute" {
		query = query + "'%H:%i'"
	}
	if timeGrain == "hour" {
		query = query + "'%m-%d %H'"
	}
	if timeGrain == "day" {
		query = query + "'%m-%d'"
	}
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("FROM_UNIXTIME( happen_time / 1000, "+query+") AS time_key, COUNT( * ) AS total, status, round( AVG( load_time ), 2 ) AS load_time").Where(sqlWhere, startTime+" 00:00:00", endTime+" 23:59:59", monitorId).Group("time_key").Find(&httpStageTime).Error
	return
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func GetHttpErrorInfo(monitorId string, startTime string, endTime string) (httpError response.PageHttpErrorResponse, err error) {
	publicSqlWhere := "from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ?"
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("Count(*) as error_400").Where(publicSqlWhere+" and status = ?", startTime+" 00:00:00", endTime+" 23:59:59", monitorId, 400).Scan(&httpError.HttpErrorQuotaResponse.Error400).Error
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("Count(*) as error_404").Where(publicSqlWhere+" and status = ?", startTime+" 00:00:00", endTime+" 23:59:59", monitorId, 404).Scan(&httpError.HttpErrorQuotaResponse.Error404).Error
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("Count(*) as error_500").Where(publicSqlWhere+" and status = ?", startTime+" 00:00:00", endTime+" 23:59:59", monitorId, 500).Scan(&httpError.HttpErrorQuotaResponse.Error500).Error
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("Count(DISTINCT user_id) as error_user").Where(publicSqlWhere+" and status = ?", startTime+" 00:00:00", endTime+" 23:59:59", monitorId, 500).Scan(&httpError.HttpErrorQuotaResponse.ErrorUser).Error
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("http_url, Count(http_url) as total, Count(DISTINCT user_id) as error_user, status, round( AVG( load_time ), 2 ) AS load_time").Where(publicSqlWhere+" And status = 400 or status = 404 or status = 500", startTime+" 00:00:00", endTime+" 23:59:59", monitorId).Group("http_url").Find(&httpError.HttpListResponse).Error
	return
}
