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
		Select("http_url as url, Count(http_url) as total, Count(DISTINCT user_id) as user_total, status, round( AVG( load_time ), 2 ) AS load_time, "+
			"(SELECT Count(DISTINCT user_id) From page_https WHERE page_https.http_url = url AND page_https.load_time > 3000) as user_slow, "+
			"round(("+
			"(SELECT Count(url) From page_https WHERE page_https.http_url = url AND page_https.status = 200)/"+
			"(SELECT Count(url) From page_https WHERE page_https.http_url = url)"+
			")*100,2) as success_rate"+
			"").Where(
		SqlWhereBuild("page_https")+" and status = 200", startTimes, endTimes, monitorId).
		Group("url").Order("load_time desc").Find(&httpInfoList).Error
	return
}

func GetHttpErrorList(monitorId string, startTime string, endTime string) (httpErrs []response.HttpErrsResponse, err error) {
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url as url, Count(http_url) as total, Count(DISTINCT user_id) as user_total, status,  request_text, "+
			"max(happen_time) as last_happen_time, min(happen_time) as first_happen_time ").Where(
		SqlWhereBuild("page_https")+" AND status != 200", startTimes, endTimes, monitorId).
		Group("url").Order("load_time desc").Find(&httpErrs).Error
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

func GetHttpStage(monitorId string, startTime string, endTime string) (httpStageTime []response.HttpStageTimeResponse, err error) {
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("FROM_UNIXTIME( happen_time / 1000, '%m-%d %H:%i') AS time_key, COUNT( * ) AS total, status, round( AVG( load_time ), 2 ) AS load_time").
		Where(SqlWhereBuild("page_https")+"and status = 200", startTime+" 00:00:00", endTime+" 23:59:59", monitorId).
		Group("time_key").Find(&httpStageTime).Error
	return
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
