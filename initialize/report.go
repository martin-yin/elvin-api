package initialize

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//

// 页面性能
func InitReportData() {
	ticker := time.NewTicker(1 * time.Second)
	var wg sync.WaitGroup
	go func() {
		for {
			select {
			case <-ticker.C:
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
							publicFiles.IP = "58.243.220.37"
							addressInfo := getIpAddressInfo(publicFiles.IP)
							publicFiles.Nation = addressInfo.Nation
							publicFiles.City = addressInfo.City
							publicFiles.District = addressInfo.District
							publicFiles.Province = addressInfo.Province
							if publicFiles.ActionType == "PAGE_LOAD" {
								var pagePerformanceBody request.PostPagePerformance
								json.Unmarshal([]byte(performance), &pagePerformanceBody)
								createPerformance(pagePerformanceBody, publicFiles)
							} else if publicFiles.ActionType == "HTTP_LOG" {
								var pageHttpBody request.PostPageHttpBody
								json.Unmarshal([]byte(performance), &pageHttpBody)
								createHttp(pageHttpBody, publicFiles)
							} else if publicFiles.ActionType == "PAGE_VIEW" {
								var pageViewBody request.PostPageViewBody
								json.Unmarshal([]byte(performance), &pageViewBody)
								createPageView(pageViewBody, publicFiles)
							} else if publicFiles.ActionType == "BEHAVIOR_INFO" {
								var behaviorInfoBody request.PostBehaviorInfoBody
								json.Unmarshal([]byte(performance), &behaviorInfoBody)
								CreatePageBehavior(behaviorInfoBody, publicFiles)
							} else if publicFiles.ActionType == "RESOURCE_ERROR" {
								var pageResourceErroBody request.PostPageResourceErroBody
								json.Unmarshal([]byte(performance), &pageResourceErroBody)
								createResourcesError(pageResourceErroBody, publicFiles)
							} else if publicFiles.ActionType == "JS_ERROR" {
								var jsErrorBody request.PostJsErrorBody
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
		}
	}()
}

func createJsError(jsErrorBody request.PostJsErrorBody, publicFiles model.PublicFiles) {
	jsErrorModel := model.PageJsError{
		PageUrl:       jsErrorBody.PageUrl,
		ComponentName: jsErrorBody.ComponentName,
		Stack:         jsErrorBody.Stack,
		Message:       jsErrorBody.Message,
		PublicFiles:   publicFiles,
	}
	if err := services.CreatePageJsError(jsErrorModel, jsErrorBody.EventId); err != nil {
		fmt.Print(err, "!!!!!!!!!")
	}
}

func createPerformance(performanceBody request.PostPagePerformance, publicFiles model.PublicFiles) {
	pagePerformanceModel := model.PagePerformance{
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
	if err := services.CreatePagePerformance(&pagePerformanceModel, publicFiles.EventId); err != nil {
		fmt.Print(err, "!!!!!!!!!!!!")
	}
}

func createHttp(pageHttpBody request.PostPageHttpBody, publicFiles model.PublicFiles) {
	HttpInfoModel := model.PageHttp{
		PageUrl:      pageHttpBody.PageUrl,
		HttpUrl:      pageHttpBody.HttpUrl,
		LoadTime:     pageHttpBody.LoadTime,
		Status:       pageHttpBody.Status,
		StatusText:   pageHttpBody.StatusText,
		StatusResult: pageHttpBody.StatusResult,
		RequestText:  pageHttpBody.RequestText,
		ResponseText: pageHttpBody.ResponseText,
		PublicFiles:  publicFiles,
	}
	if err := services.CreatePageHttpModel(HttpInfoModel, pageHttpBody.EventId); err != nil {
		fmt.Print(err, "!!!!!!!!!!!!")
	}
}

func createResourcesError(pageResourceErroBody request.PostPageResourceErroBody, publicFiles model.PublicFiles) {
	resourceErrorInfoModel := model.PageResourceError{
		PageUrl:     pageResourceErroBody.PageUrl,
		SourceUrl:   pageResourceErroBody.SourceUrl,
		ElementType: pageResourceErroBody.ElementType,
		Status:      pageResourceErroBody.Status,
		PublicFiles: publicFiles,
	}
	if err := services.CreateResourcesError(resourceErrorInfoModel, pageResourceErroBody.EventId); err != nil {
		fmt.Print(err, "!!!!!!!!!!!!")
	}
}

func createPageView(pageViewBody request.PostPageViewBody, publicFiles model.PublicFiles) {
	pageViewModel := model.PageView{
		PageUrl:     pageViewBody.PageUrl,
		PublicFiles: publicFiles,
	}
	if err := services.CreatePageView(pageViewModel, pageViewBody.EventId); err != nil {
		fmt.Print(err, "!!!!!!!!!!!!")
	}
}

func CreatePageBehavior(behaviorInfoBody request.PostBehaviorInfoBody, publicFiles model.PublicFiles) {
	pageBehaviorInfoModel := model.PageBehavior{
		PageUrl:     behaviorInfoBody.PageUrl,
		ClassName:   behaviorInfoBody.ClassName,
		Placeholder: behaviorInfoBody.Placeholder,
		InputValue:  behaviorInfoBody.InputValue,
		TagNameint:  behaviorInfoBody.TagNameint,
		InnterText:  behaviorInfoBody.InnterText,
		PublicFiles: publicFiles,
	}
	if err := services.CreatePageBehavior(pageBehaviorInfoModel, behaviorInfoBody.EventId); err != nil {
		fmt.Print(err, "!!!!!!!!!!!!")
	}
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
