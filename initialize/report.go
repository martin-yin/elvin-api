package initialize

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/services"
	"encoding/json"
	"time"
)

func InitReportData()  {
	// 这里开始跑定时任务去写数据
	reportPerformanceData();
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
	var pagePerformanceBody model.PagePerformance
	var pagePerformanceBodyList []model.PagePerformance
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				pipe := global.GVA_REDIS.TxPipeline()
				getValue := pipe.LRange("performance", 0, 1000)
				_, err := pipe.Exec()
				if err !=nil {
					return
				}
				if len(getValue.Val()) > 0 {
					for _, value := range getValue.Val() {
						err = json.Unmarshal([]byte(value), &pagePerformanceBody)
						pagePerformanceBodyList = append(pagePerformanceBodyList, model.PagePerformance{
							PageUrl:        pagePerformanceBody.PageUrl,
							Appcache:       pagePerformanceBody.Appcache,
							LookupDomain:   pagePerformanceBody.LookupDomain,
							Tcp:            pagePerformanceBody.Tcp,
							SslT:           pagePerformanceBody.SslT,
							Request:        pagePerformanceBody.Request,
							DomParse:       pagePerformanceBody.DomParse,
							Ttfb:           pagePerformanceBody.Ttfb,
							LoadPage:       pagePerformanceBody.LoadPage,
							LoadEvent:      pagePerformanceBody.LoadEvent,
							LoadType:       pagePerformanceBody.LoadType,
							Redirect:       pagePerformanceBody.Redirect,
							PublicFiles: pagePerformanceBody.PublicFiles,
						})
					}
					if err := services.CreatePagePerformance(&pagePerformanceBodyList); err != nil {
						return
					}
					pagePerformanceBodyList = nil
					global.GVA_REDIS.LTrim("performance", 500, -1)
				}
			}
		}
	}()
}
