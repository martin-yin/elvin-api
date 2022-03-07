package initialize

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/services"
	"dancin-api/utils"
	"encoding/json"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"sync"
)

var handles *utils.Handles

func consume(report string) {
	var commonFiles model.CommonFiles
	json.Unmarshal([]byte(report), &commonFiles)
	// 向chan 添加ip地址
	utils.IPChan <- commonFiles.IP
	services.CreateUserAction(commonFiles, report)
	handles.ServiceHandlers[commonFiles.ActionType](report, &commonFiles)
}

// kafka 消费数据
func ReportDataConsumeByKafka() {
	for {
		m, err := global.KAFKA.ReadMessage(100)
		if err != nil {
			global.LOGGER.Error("kafka读取数据失败:", zap.Any("err", err))
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

func consumeByRedis() {
	wg := sync.WaitGroup{}
	pipe := global.REDIS.TxPipeline()
	reportList := pipe.LRange("reportData", 0, 10000)
	_, err := pipe.Exec()
	if err != nil {
		global.LOGGER.Error("redis 数据读取出错:", zap.Any("err", err))
		return
	}
	if len(reportList.Val()) > 0 {
		for _, report := range reportList.Val() {
			wg.Add(1)
			go func(wg *sync.WaitGroup, report string) {
				consume(report)
				defer wg.Done()
			}(&wg, report)
		}
		wg.Wait()
		global.REDIS.LTrim("reportData", 10000, -1)
	}
}

func init() {
	handles = utils.NewHandles()
	servicesHandles := map[string]utils.ServiceFunc{
		"PERFORMANCE": func(report string, commonFiles *model.CommonFiles) {
			var performance request.PerformanceBody
			json.Unmarshal([]byte(report), &performance)
			services.CreatePagePerformance(&performance, commonFiles)
		},
		"HTTPLOG": func(report string, commonFiles *model.CommonFiles) {
			var http request.HttpBody
			json.Unmarshal([]byte(report), &http)
			services.CreatePageHttp(&http, commonFiles)
		},
		"PAGEVIEW": func(report string, commonFiles *model.CommonFiles) {
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
		"JSERROR": func(report string, commonFiles *model.CommonFiles) {
			var issuesBody request.IssuesBody
			json.Unmarshal([]byte(report), &issuesBody)
			services.CreatePageJsError(&issuesBody, commonFiles)
		},
	}
	handles.ServicesHandlerRegister(servicesHandles)
	go utils.ConsumeIP()
}
