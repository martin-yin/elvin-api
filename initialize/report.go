package initialize

import (
	"context"
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var handles *utils.Handles

func reportDataConsume() {
	for {
		m, err := global.GVA_KAFKA.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		var publicFiles model.PublicFiles
		report := string(m.Value)
		json.Unmarshal([]byte(report), &publicFiles)
		addressInfo := getIpAddressInfo(publicFiles.IP)
		publicFiles.Nation = addressInfo.Nation
		publicFiles.City = addressInfo.City
		publicFiles.District = addressInfo.District
		publicFiles.Province = addressInfo.Province
		services.CreateUserAction(publicFiles, report)
		handles.ServiceHandlers[publicFiles.ActionType](report, &publicFiles)
	}
}

//func reportDataConsume() {
//	var wg sync.WaitGroup
//	pipe := global.GVA_REDIS.TxPipeline()
//	reportList := pipe.LRange("reportData", 0, 10000)
//	_, err := pipe.Exec()
//	if err != nil {
//		return
//	}
//	if len(reportList.Val()) > 0 {
//		for _, report := range reportList.Val() {
//			wg.Add(1)
//			go func(report string) {
//				var publicFiles model.PublicFiles
//				json.Unmarshal([]byte(report), &publicFiles)
//				addressInfo := getIpAddressInfo(publicFiles.IP)
//				publicFiles.Nation = addressInfo.Nation
//				publicFiles.City = addressInfo.City
//				publicFiles.District = addressInfo.District
//				publicFiles.Province = addressInfo.Province
//				services.CreateUserAction(publicFiles, report)
//				handles.ServiceHandlers[publicFiles.ActionType](report, &publicFiles)
//				wg.Done()
//			}(report)
//		}
//		wg.Wait()
//		global.GVA_REDIS.LTrim("reportData", 10000, -1)
//	}
//}

func init() {
	handles = utils.NewHandles()
	servicesHandles := map[string]utils.ServiceFunc{
		"PAGE_LOAD": func(report string, publicFiles *model.PublicFiles) {
			var performance request.PerformanceBody
			json.Unmarshal([]byte(report), &performance)
			services.CreatePagePerformance(&performance, publicFiles)
		},
		"HTTP_LOG": func(report string, publicFiles *model.PublicFiles) {
			var http request.HttpBody
			json.Unmarshal([]byte(report), &http)
			services.CreatePageHttp(&http, publicFiles)
		},
		"PAGE_VIEW": func(report string, publicFiles *model.PublicFiles) {
			var pageView request.PageViewBody
			json.Unmarshal([]byte(report), &pageView)
			services.CreatePageView(&pageView, publicFiles)
		},
		"OPERATION": func(report string, publicFiles *model.PublicFiles) {
			var operation request.OperationBody
			json.Unmarshal([]byte(report), &operation)
			services.CreatePageOperation(&operation, publicFiles)
		},
		"RESOURCE": func(report string, publicFiles *model.PublicFiles) {
			var resource request.ResourceErrorBody
			json.Unmarshal([]byte(report), &resource)
			services.CreateResourcesError(&resource, publicFiles)
		},
		"JS_ERROR": func(report string, publicFiles *model.PublicFiles) {
			var issuesBody request.IssuesBody
			json.Unmarshal([]byte(report), &issuesBody)
			services.CreatePageJsError(&issuesBody, publicFiles)
		},
	}
	handles.ServicesHandlerRegister(servicesHandles)
}
func InitReportData() {
	KafkaReader()
	reportDataConsume()
}

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

func getIpAddressInfo(ip string) (AdInfo response.TxMapResultAdInfo) {
	if ip == "" {
		return
	}
	var txMapResponse response.TxMapResponse
	addingStr := global.GVA_REDIS.HGet("ipAddress", ip)
	if len(addingStr.Val()) != 0 {
		err := json.Unmarshal([]byte(addingStr.Val()), &AdInfo)
		if err != nil {
			fmt.Print(err, "出粗了！")
		}
		return AdInfo
	} else {
		resp, err := http.Get("https://apis.map.qq.com/ws/location/v1/ip?ip=" + ip + "&key=TFNBZ-STIKX-JQ242-TNUNK-4NWCT-CLF7S")
		if err != nil {
			return
		}
		txMapResponded, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		err = json.Unmarshal(txMapResponded, &txMapResponse)
		txMapResponseStr, err := json.Marshal(&txMapResponse.Result.AdInfo)
		global.GVA_REDIS.HSet("ipAddress", ip, txMapResponseStr)
		return txMapResponse.Result.AdInfo
	}
}
