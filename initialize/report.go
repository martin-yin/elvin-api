package initialize

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"danci-api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//
func InitReportData() {
	// 这里开始跑定时任务去写数据
	reportPerformanceData()
}

// http请求写入
func reportHttpData() {

}

// pageJsError
func reportPageJsError() {
}

// 资源报错
func reportResourceError() {
}

// 点击事件
func reportPageBehavior() {
}

func reportPerformanceData() {
	//var pagePerformanceBodyList []model.PagePerformance
	ticker := time.NewTicker(1 * time.Second)
	performanceChannel := make(chan string)
	//var wg sync.WaitGroup
	go func() {
		for {
			select {
			case <-ticker.C:
				pipe := global.GVA_REDIS.TxPipeline()
				getValue := pipe.LRange("performance", 0, 1)
				_, err := pipe.Exec()
				if err != nil {
					return
				}
				if len(getValue.Val()) > 0 {
					for _, value := range getValue.Val() {
						// 向channel 中写入数据。
						performanceChannel <- value
					}
					global.GVA_REDIS.LTrim("performance", 1, -1)
				}
			}
		}
	}()

	go func() {
		for performance := range performanceChannel {
			var pagePerformanceBody model.PagePerformance
			json.Unmarshal([]byte(performance), &pagePerformanceBody)
			addressInfo := GetIpAddressInfo(pagePerformanceBody.PublicFiles.IP)
			pagePerformanceBody.PublicFiles.Nation = addressInfo.Nation
			pagePerformanceBody.PublicFiles.City = addressInfo.City
			pagePerformanceBody.PublicFiles.District = addressInfo.District
			pagePerformanceBody.PublicFiles.Province = addressInfo.Province
			pagePerformanceModel := model.PagePerformance{
				PageUrl:      pagePerformanceBody.PageUrl,
				Appcache:     pagePerformanceBody.Appcache,
				LookupDomain: pagePerformanceBody.LookupDomain,
				Tcp:          pagePerformanceBody.Tcp,
				SslT:         pagePerformanceBody.SslT,
				Request:      pagePerformanceBody.Request,
				DomParse:     pagePerformanceBody.DomParse,
				Ttfb:         pagePerformanceBody.Ttfb,
				LoadPage:     pagePerformanceBody.LoadPage,
				LoadEvent:    pagePerformanceBody.LoadEvent,
				LoadType:     pagePerformanceBody.LoadType,
				Redirect:     pagePerformanceBody.Redirect,
				PublicFiles:  pagePerformanceBody.PublicFiles,
			}
			if err := services.CreatePagePerformance(&pagePerformanceModel); err != nil {
				fmt.Print(err, "!!!!!!!!!!!!")
			}
		}
	}()
}

func GetIpAddressInfo(ip string) (AdInfo response.TxMapResultAdInfo) {
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
		txMapbody, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		err = json.Unmarshal(txMapbody, &txMapResponse)
		txMapResponseStr, err := json.Marshal(&txMapResponse.Result.AdInfo)
		global.GVA_REDIS.HSet("ipAddress", ip, txMapResponseStr)
		return txMapResponse.Result.AdInfo
	}
}
