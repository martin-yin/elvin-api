package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"fmt"
	"strconv"
)

func GetHttpInfoList(params request.RequestParams) (httpInfoList []response.HttpListResponse, err error) {
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url as url, Count(http_url) as total, Count(DISTINCT user_id) as user_total, status, round( AVG( load_time ), 2 ) AS load_time, "+
			"(SELECT Count(DISTINCT user_id) From page_https WHERE page_https.http_url = url AND page_https.load_time > 3000) as user_slow, "+
			"round(("+
			"(SELECT Count(url) From page_https WHERE page_https.http_url = url AND page_https.status = 200)/"+
			"(SELECT Count(url) From page_https WHERE page_https.http_url = url)"+
			")*100,2) as success_rate"+
			"").Where(
		SqlWhereBuild("page_https")+" and status = 200", params.StartTime, params.EndTime, params.MonitorId).
		Group("url").Order("load_time desc").Find(&httpInfoList).Error
	return
}

func GetHttpErrorList(params request.RequestParams) (httpErrs []response.HttpErrsResponse, err error) {
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url as url, Count(http_url) as total, Count(DISTINCT user_id) as user_total, status,  request_text, "+
			"max(happen_time) as last_happen_time, min(happen_time) as first_happen_time ").Where(
		SqlWhereBuild("page_https")+" AND status != 200", params.StartTime, params.EndTime, params.MonitorId).
		Group("url").Order("load_time desc").Find(&httpErrs).Error
	return
}

func GetHttpQuota(params request.RequestParams) (httpQuota response.HttpQuotaResponse, err error) {
	err = global.GORMDB.Model(&model.PageHttp{}).Select("round(AVG( load_time )) AS load_time, status, COUNT(*) AS total").Where(
		SqlWhereBuild("page_https"),
		params.StartTime, params.EndTime, params.MonitorId).Find(&httpQuota).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select(" Count(DISTINCT user_id) as error_user, status ").
		Where(SqlWhereBuild("page_https")+" and status > 400", params.StartTime, params.EndTime, params.MonitorId).Find(&httpQuota).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select("COUNT(*) AS success_total, status").
		Where(SqlWhereBuild("page_https")+" and status < 400", params.StartTime, params.EndTime, params.MonitorId).Scan(&httpQuota).Error
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
