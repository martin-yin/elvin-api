package initialize

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"sync"
)

var handles *utils.Handles

func reportDataConsume() {
	var wg sync.WaitGroup
	pipe := global.GVA_REDIS.TxPipeline()
	getValue := pipe.LRange("reportData", 0, 5000)
	_, err := pipe.Exec()
	if err != nil {
		return
	}
	if len(getValue.Val()) > 0 {
		for _, performance := range getValue.Val() {
			wg.Add(1)
			go func(performance string) {
				reportConsumer(performance)
			}(performance)
		}
		global.GVA_REDIS.LTrim("reportData", 5000, -1)
		wg.Wait()
	}
}

func reportConsumer(report string) {
	var publicFiles model.PublicFiles
	json.Unmarshal([]byte(report), &publicFiles)
	addressInfo := getIpAddressInfo(publicFiles.IP)
	publicFiles.Nation = addressInfo.Nation
	publicFiles.City = addressInfo.City
	publicFiles.District = addressInfo.District
	publicFiles.Province = addressInfo.Province
	services.CreateUserAction(publicFiles, report)
	handles.ServiceHandlers[publicFiles.ActionType](report, &publicFiles)
}


func init() {
	handles = utils.NewHandles()
	handles.ServicesHandlerRegister("PAGE_LOAD", func(report string, publicFiles *model.PublicFiles) {
		var performance request.PerformanceBody
		json.Unmarshal([]byte(report), &performance)
		services.CreatePagePerformance(&performance, publicFiles)
	})

	handles.ServicesHandlerRegister("HTTP_LOG", func(report string, publicFiles *model.PublicFiles) {
		var http request.HttpBody
		json.Unmarshal([]byte(report), &http)
		services.CreatePageHttp(&http, publicFiles)
	})

	handles.ServicesHandlerRegister("PAGE_VIEW", func(report string, publicFiles *model.PublicFiles) {
		var pageView request.PageViewBody
		json.Unmarshal([]byte(report), &pageView)
		services.CreatePageView(&pageView, publicFiles)
	})

	handles.ServicesHandlerRegister("OPERATION", func(report string, publicFiles *model.PublicFiles) {
		var operation request.OperationBody
		json.Unmarshal([]byte(report), &operation)
		services.CreatePageOperation(&operation, publicFiles)
	})

	handles.ServicesHandlerRegister("RESOURCE", func(report string, publicFiles *model.PublicFiles) {
		var resource request.ResourceErrorBody
		json.Unmarshal([]byte(report), &resource)
		services.CreateResourcesError(&resource, publicFiles)
	})

	handles.ServicesHandlerRegister("JS_ERROR", func(report string, publicFiles *model.PublicFiles) {
		var issuesBody request.IssuesBody
		json.Unmarshal([]byte(report), &issuesBody)
		services.CreatePageJsError(&issuesBody, publicFiles)
	})
}

// 页面性能

func InitReportData() {

	cron2 := cron.New(cron.WithSeconds())
	cron2.AddFunc("*/5 * * * * * ", reportDataConsume)
	//cron2.AddFunc("0 0 0 1 * ?  ", func() {   这个是正式得，每天凌晨调用一次。
	//cron2.AddFunc("*/10 * * * * * ", func() {  // 这个是测试时使用的。
	//
	//	actionType := [7]string{"PAGE_LOAD", "HTTP_ERROR_LOG", "HTTP_LOG" , "RESOURCE_ERROR", "BEHAVIOR_INFO", "PAGE_VIEW", "JS_ERROR"}
	//	var reportData []model.ReportDayStatistic
	//	startTime := time.Now().Format("2006-01-02")
	//	for _, value := range projectList {
	//		for _, action :=range actionType {
	//			keyName := startTime + value.MonitorId + action;
	//			count := global.GVA_REDIS.Get(keyName).Val()
	//			if count != "" {
	//				reportData = append(reportData, model.ReportDayStatistic{
	//					ActionType: "PAGE_LOAD",
	//					MonitorId: value.MonitorId,
	//					Day: startTime,
	//					Count: global.GVA_REDIS.Get(keyName).Val(),
	//				})
	//				global.GVA_REDIS.Del(keyName)
	//			}
	//		}
	//	}
	//	services.CreateReportDay(reportData)
	//})
	cron2.Start()
}

func getIpAddressInfo(ip string) (AdInfo response.TxMapResultAdInfo) {
	if ip == "" {
		return
	}
	var txMapResponse response.TxMapResponse
	adinfoStr := global.GVA_REDIS.HGet("ipAddress", ip)
	if len(adinfoStr.Val()) != 0 {
		err := json.Unmarshal([]byte(adinfoStr.Val()), &AdInfo)
		if err != nil {
			fmt.Print(err, "出粗了！")
		}
		return AdInfo
	} else {
		resp, err := http.Get("https://apis.map.qq.com/ws/location/v1/ip?ip=" + ip + "&key=TFNBZ-STIKX-JQ242-TNUNK-4NWCT-CLF7S")
		if err != nil {
			return
		}
		txMapResponsebody, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		err = json.Unmarshal(txMapResponsebody, &txMapResponse)
		txMapResponseStr, err := json.Marshal(&txMapResponse.Result.AdInfo)
		global.GVA_REDIS.HSet("ipAddress", ip, txMapResponseStr)
		return txMapResponse.Result.AdInfo
	}
}
