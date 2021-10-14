package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/response"
	"fmt"
	"strconv"
)

func GetHttpInfoList(monitorId string, startTime string, endTime string) (httpInfoList []response.HttpListResponse, err error) {
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url as url, Count(DISTINCT http_url) as total, Count(DISTINCT user_id) as user_total, status, round( AVG( load_time ), 2 ) AS load_time, "+
			"(SELECT Count(DISTINCT user_id) From page_https WHERE page_https.http_url = url AND page_https.load_time > 3000) as user_slow ").Where(
		SqlWhereBuild("page_https")+" and status = 200", startTimes, endTimes, monitorId).
		Group("url").Order("load_time desc").Find(&httpInfoList).Error
	return
}

func GetHttpErrorList(monitorId string, startTime string, endTime string) (httpInfoList []response.HttpListResponse, err error) {
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url, Count(DISTINCT http_url) as total, Count(DISTINCT user_id) as user_total, status, round( AVG( load_time ), 2 ) AS load_time").Where(
		SqlWhereBuild("page_https")+" and status != 200", startTimes, endTimes, monitorId).
		Group("http_url").Order("load_time desc").Find(&httpInfoList).Error
	return
}

func GetHttpQuota(monitorId string, startTime string, endTime string) (httpQuota response.HttpQuotaResponse, err error) {
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GORMDB.Model(&model.PageHttp{}).Select("round(AVG( load_time )) AS load_time, status, COUNT(*) AS total").Where(
		SqlWhereBuild("page_https"),
		startTimes, endTimes, monitorId).Find(&httpQuota).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select(" Count(DISTINCT user_id) as error_user, status ").
		Where(SqlWhereBuild("page_https")+" and status > 400", startTimes, endTimes, monitorId).Find(&httpQuota).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select("COUNT(*) AS success_total, status").
		Where(SqlWhereBuild("page_https")+" and status < 400", startTimes, endTimes, monitorId).Scan(&httpQuota).Error
	return
}

func GetHttpStage(monitorId string, timeGrain string, startTime string, endTime string) (httpStageTime []response.HttpStageTimeResponse, err error) {
	query := ""
	if timeGrain == "minute" {
		query = query + "'%H:%i'"
	}
	if timeGrain == "hour" {
		query = query + "'%m-%d %H'"
	}
	if timeGrain == "day" {
		query = query + "'%m-%d'"
	}
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("FROM_UNIXTIME( happen_time / 1000, "+query+") AS time_key, COUNT( * ) AS total, status, round( AVG( load_time ), 2 ) AS load_time").
		Where(SqlWhereBuild("page_https")+"and status = 200", startTime+" 00:00:00", endTime+" 23:59:59", monitorId).
		Group("time_key").Find(&httpStageTime).Error
	return
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
