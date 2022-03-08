package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/utils"
	"fmt"
	"strconv"
)

func GetHttpList(params request.RequestParams) (httpInfoList []response.HttpListResponse, err error) {
	sql, sqlParams := utils.BuildWhereSql("page_https", " and status = 200", params)
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url as url, Count(http_url) as total, Count(DISTINCT user_id) as user_total, status, round( AVG( load_time ), 2 ) AS load_time, "+
			"(SELECT Count(DISTINCT user_id) From page_https WHERE page_https.http_url = url AND page_https.load_time > 3000) as user_slow, "+
			"round(("+
			"(SELECT Count(url) From page_https WHERE page_https.http_url = url AND page_https.status = 200)/"+
			"(SELECT Count(url) From page_https WHERE page_https.http_url = url)"+
			")*100,2) as success_rate"+
			"").Where(sql, sqlParams...).
		Group("url").Order("load_time desc").Find(&httpInfoList).Error
	return
}

func GetHttpErrorList(params request.RequestParams) (httpErrs []response.HttpErrsResponse, err error) {
	sql, sqlParams := utils.BuildWhereSql("page_https", " and status != 200", params)
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("http_url as url, Count(http_url) as total, Count(DISTINCT user_id) as user_total, status,  request_text, "+
			"max(happen_time) as last_happen_time, min(happen_time) as first_happen_time ").Where(sql, sqlParams...).
		Group("url").Order("load_time desc").Find(&httpErrs).Error
	return
}

func GetHttpQuota(params request.RequestParams) (httpQuota response.HttpQuotaResponse, err error) {
	sql, sqlParams := utils.BuildWhereSql("page_https", " and status = 200", params)
	err = global.GORMDB.Model(&model.PageHttp{}).Select("round(AVG( load_time )) AS load_time, status, COUNT(*) AS total").Where(sql, sqlParams...).Find(&httpQuota).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select(" Count(DISTINCT user_id) as error_user, status ").Where(sql+" and status > 400", sqlParams...).Find(&httpQuota).Error
	err = global.GORMDB.Model(&model.PageHttp{}).Select("COUNT(*) AS success_total, status").
		Where(sql+" and status < 400", sqlParams...).Scan(&httpQuota).Error
	return
}

func GetHttpStage(params request.RequestParams) (httpStageTime []response.HttpStageTimeResponse, err error) {
	sql, sqlParams := utils.BuildWhereSql("page_https", " and status = 200", params)
	err = global.GORMDB.Model(&model.PageHttp{}).
		Select("FROM_UNIXTIME( happen_time / 1000, '%m-%d %H:%i') AS time_key, COUNT( * ) AS total, status, round( AVG( load_time ), 2 ) AS load_time").
		Where(sql, sqlParams...).Group("time_key").Find(&httpStageTime).Error
	return
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
