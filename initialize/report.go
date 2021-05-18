package initialize

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"sync"
)

func reportDataWrite() {
	var wg sync.WaitGroup
	pipe := global.GVA_REDIS.TxPipeline()
	getValue := pipe.LRange("reportData", 0, 500)
	_, err := pipe.Exec()
	if err != nil {
		return
	}
	if len(getValue.Val()) > 0 {
		for _, performance := range getValue.Val() {
			wg.Add(1)
			go func(performance string) {
				var publicFiles model.PublicFiles
				json.Unmarshal([]byte(performance), &publicFiles)
				addressInfo := getIpAddressInfo(publicFiles.IP)
				publicFiles.Nation = addressInfo.Nation
				publicFiles.City = addressInfo.City
				publicFiles.District = addressInfo.District
				publicFiles.Province = addressInfo.Province
				services.CreateUserAction(publicFiles, performance)
				if publicFiles.ActionType == "PAGE_LOAD" {
					var pagePerformanceBody request.PerformanceBody
					json.Unmarshal([]byte(performance), &pagePerformanceBody)
					createPerformance(pagePerformanceBody, publicFiles)
				} else if publicFiles.ActionType == "HTTP_LOG" {
					var pageHttpBody request.HttpBody
					json.Unmarshal([]byte(performance), &pageHttpBody)
					createHttp(pageHttpBody, publicFiles)
				} else if publicFiles.ActionType == "PAGE_VIEW" {
					var pageViewBody request.PageViewBody
					json.Unmarshal([]byte(performance), &pageViewBody)
					createPageView(pageViewBody, publicFiles)
				} else if publicFiles.ActionType == "BEHAVIOR_INFO" {
					var operationBody request.OperationBody
					json.Unmarshal([]byte(performance), &operationBody)
					CreatePageBehavior(operationBody, publicFiles)
				} else if publicFiles.ActionType == "RESOURCE_ERROR" {
					var resourceErroBody request.ResourceErrorBody
					json.Unmarshal([]byte(performance), &resourceErroBody)
					createResourcesError(resourceErroBody, publicFiles)
				} else if publicFiles.ActionType == "JS_ERROR" {
					var jsErrorBody request.JsErrorBody
					json.Unmarshal([]byte(performance), &jsErrorBody)
					createJsError(jsErrorBody, publicFiles)
				}
				wg.Done()
			}(performance)
		}
		wg.Wait()
		global.GVA_REDIS.LTrim("reportData", 500, -1)
	}
}

// 页面性能
func InitReportData() {
	cron2 := cron.New(cron.WithSeconds())
	cron2.AddFunc("*/10 * * * * * ", reportDataWrite)
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
	//cron2.Start()
	cron2.Start()
}

func createJsError(jsErrorBody request.JsErrorBody, publicFiles model.PublicFiles) {
	jsError := model.PageJsError{
		PageUrl:       jsErrorBody.PageUrl,
		ComponentName: jsErrorBody.ComponentName,
		Stack:         jsErrorBody.Stack,
		Message:       jsErrorBody.Message,
		StackFrames:   jsErrorBody.StackFrames,
		ErrorName:     jsErrorBody.ErrorName,
		PublicFiles:   publicFiles,
	}
	jsIssueModel, err := services.FindJsIssue(jsErrorBody.Message)
	if err == nil {
		if jsIssueModel.ID != 0 {
			jsError.JsIssuesId = jsIssueModel.ID
			services.CreatePageJsError(jsError)
		} else {
			jsIssue := model.JsIssue{
				ErrorName: jsError.ErrorName,
				Message:   jsError.Message,
				PageJsError: []model.PageJsError{
					jsError,
				},
			}
			services.CreateJsIssue(jsIssue)
		}
	}
}

func createPerformance(performanceBody request.PerformanceBody, publicFiles model.PublicFiles) {
	performance := model.PagePerformance{
		PageUrl:      performanceBody.PageUrl,
		Appcache:     performanceBody.Appcache,
		LookupDomain: performanceBody.LookupDomain,
		Tcp:          performanceBody.Tcp,
		SslT:         performanceBody.SslT,
		Request:      performanceBody.Request,
		DomParse:     performanceBody.DomParse,
		Ttfb:         performanceBody.Ttfb,
		LoadPage:     performanceBody.LoadPage,
		LoadEvent:    performanceBody.LoadEvent,
		LoadType:     performanceBody.LoadType,
		Redirect:     performanceBody.Redirect,
		PublicFiles:  publicFiles,
	}
	services.CreatePagePerformance(&performance)
}

func createHttp(httpBody request.HttpBody, publicFiles model.PublicFiles) {
	http := model.PageHttp{
		PageUrl:      httpBody.PageUrl,
		HttpUrl:      httpBody.HttpUrl,
		LoadTime:     httpBody.LoadTime,
		Method:       httpBody.Method,
		Status:       httpBody.Status,
		StatusText:   httpBody.StatusText,
		StatusResult: httpBody.StatusResult,
		RequestText:  httpBody.RequestText,
		ResponseText: httpBody.ResponseText,
		PublicFiles:  publicFiles,
	}
	services.CreatePageHttp(&http)

}

func createResourcesError(resourceErrorBody request.ResourceErrorBody, publicFiles model.PublicFiles) {
	resourceError := model.PageResourceError{
		PageUrl:     resourceErrorBody.PageUrl,
		SourceUrl:   resourceErrorBody.SourceUrl,
		ElementType: resourceErrorBody.ElementType,
		Status:      resourceErrorBody.Status,
		PublicFiles: publicFiles,
	}
	services.CreateResourcesError(&resourceError)

}

func createPageView(pageViewBody request.PageViewBody, publicFiles model.PublicFiles) {
	pageView := model.PageView{
		PageUrl:     pageViewBody.PageUrl,
		PublicFiles: publicFiles,
	}
	services.CreatePageView(&pageView)

}

func CreatePageBehavior(operationBody request.OperationBody, publicFiles model.PublicFiles) {
	operation := model.PageOperation{
		PageUrl:     operationBody.PageUrl,
		ClassName:   operationBody.ClassName,
		Placeholder: operationBody.Placeholder,
		InputValue:  operationBody.InputValue,
		TagNameint:  operationBody.TagNameint,
		InnterText:  operationBody.InnterText,
		PublicFiles: publicFiles,
	}
	services.CreatePageBehavior(&operation)
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
