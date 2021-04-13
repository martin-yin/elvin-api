package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func GetStackPerformance(startTime string, endTime string) (stackData response.StackResponse, err error) {
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("round(AVG(redirect),2) as redirect, "+
		"round(AVG(appcache),2) as appcache, "+
		"round(AVG(lookup_domain),2) as lookup_domain, "+
		"round(AVG(tcp),2) as tcp, "+
		"round(AVG(ssl_t),2) as ssl_t, "+
		"round(AVG(request),2) as request, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_event),2) as load_event, "+
		"round(AVG(dom_parse),2) as dom_parse ").Where("from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Find(&stackData).Error
	return
}

func GetQuotaData(startTime string, endTime string) (quotaData response.QuotaResponse, err error) {
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("round(AVG(dom_parse),2) as dom_parse, "+
		"round(AVG(ttfb),2) as ttfb, "+
		"round(AVG(load_page),2) as load_page, "+
		"Count(*) as Pv ").Where("from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quotaData).Error
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("COUNT( * ) AS fast").Where("from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&quotaData.Fast).Error

	if quotaData.Fast != 0 {
		quotaData.Fast  = Decimal(quotaData.Fast/quotaData.Pv) * 100
	}
	return
}

func GetStageTimeList(startTime string, endTime string, timeGrain string) (stageTimeList []response.StageTimeResponse, err error) {
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

	fmt.Print(query)
	fmt.Print(startTime, "start_time \n")
	fmt.Print(endTime, "endTime \n")

	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("from_unixtime(happen_time / 1000, "+query+") AS time_key, "+
		"round( AVG( redirect ), 2 ) AS redirect,"+
		"round( AVG( appcache ), 2 ) AS appcache,"+
		"round( AVG( lookup_domain ), 2 ) AS lookup_domain,"+
		"round( AVG( tcp ), 2 ) AS tcp,"+
		"round( AVG( ssl_t ), 2 ) AS ssl_t,"+
		"round( AVG( request ), 2 ) AS request,"+
		"round( AVG( ttfb ), 2 ) AS ttfb,"+
		"round( AVG( load_event ), 2 ) AS load_event,"+
		"round( AVG( load_page ), 2 ) AS load_page,"+
		"COUNT(*) as Pv ").Where("from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("time_key").Scan(&stageTimeList).Error
	return
}

func GetLoadInfoPageList(startTime string, endTime string) (pagePerformanceList []response.PagePerformanceListResponse, err error) {
	err = global.GVA_DB.Model(&model.PagePerformance{}).Select("ID, page_url, "+
		"request, "+
		"dom_parse, "+
		"ttfb, "+
		"load_page, "+
		"load_event, "+
		"load_type, "+
		"load_page, "+
		"COUNT(*) as pv ").Where("from_unixtime(page_performances.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Group("page_url").Scan(&pagePerformanceList).Error
	return
}
