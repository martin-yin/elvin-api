package initialize

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"dancin-api/utils"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"sync"
)

var handles *utils.Handles

func consume(report string) {
	var commonFiles model.CommonFiles
	json.Unmarshal([]byte(report), &commonFiles)
	addressInfo := getIpAddressInfo(commonFiles.IP)
	commonFiles.Nation = addressInfo.Nation
	commonFiles.City = addressInfo.City
	commonFiles.District = addressInfo.District
	commonFiles.Province = addressInfo.Province
	services.CreateUserAction(commonFiles, report)
	handles.ServiceHandlers[commonFiles.ActionType](report, &commonFiles)
}

// reportDataConsumeByKafka redis 消费数据
func ReportDataConsumeByKafka() {
	for {
		m, err := global.KAFKA.ReadMessage(100)
		if err != nil {
			global.LOGGER.Error("读取数据失败！！！！！！:", zap.Any("err", err))
			return
		}
		report := string(m.Value)
		consume(report)
	}
}

func ReportDataConsumeByRedis() {
	cron2 := cron.New(cron.WithSeconds())
	cron2.AddFunc("*/5 * * * * * ", consumeByRedis)
	cron2.Start()
}

func consumeByRedis(){
	var wg sync.WaitGroup
	pipe := global.REDIS.TxPipeline()
	reportList := pipe.LRange("reportData", 0, 10000)
	_, err := pipe.Exec()
	if err != nil {
		fmt.Print("fuck 。", err)
		return
	}
	if len(reportList.Val()) > 0 {
		for _, report := range reportList.Val() {
			wg.Add(1)
			go func(report string) {
				consume(report)
				wg.Done()
			}(report)
		}
		wg.Wait()
		global.REDIS.LTrim("reportData", 10000, -1)
	}
}




func init() {
	handles = utils.NewHandles()
	servicesHandles := map[string]utils.ServiceFunc{
		"PAGE_LOAD": func(report string, commonFiles *model.CommonFiles) {
			var performance request.PerformanceBody
			json.Unmarshal([]byte(report), &performance)
			services.CreatePagePerformance(&performance, commonFiles)
		},
		"HTTP_LOG": func(report string, commonFiles *model.CommonFiles) {
			var http request.HttpBody
			json.Unmarshal([]byte(report), &http)
			services.CreatePageHttp(&http, commonFiles)
		},
		"PAGE_VIEW": func(report string, commonFiles *model.CommonFiles) {
			var pageView request.PageViewBody
			json.Unmarshal([]byte(report), &pageView)
			services.CreatePageView(&pageView, commonFiles)
		},
		"OPERATION": func(report string, commonFiles *model.CommonFiles) {
			var operation request.OperationBody
			json.Unmarshal([]byte(report), &operation)
			services.CreatePageOperation(&operation, commonFiles)
		},
		"RESOURCE": func(report string, commonFiles *model.CommonFiles) {
			var resource request.ResourceErrorBody
			json.Unmarshal([]byte(report), &resource)
			services.CreateResourcesError(&resource, commonFiles)
		},
		"JS_ERROR": func(report string, commonFiles *model.CommonFiles) {
			var issuesBody request.IssuesBody
			json.Unmarshal([]byte(report), &issuesBody)
			services.CreatePageJsError(&issuesBody, commonFiles)
		},
	}
	handles.ServicesHandlerRegister(servicesHandles)
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
	addingStr := global.REDIS.HGet("ipAddress", ip)
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
		global.REDIS.HSet("ipAddress", ip, txMapResponseStr)
		return txMapResponse.Result.AdInfo
	}
}
