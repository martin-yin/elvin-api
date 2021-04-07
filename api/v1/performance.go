package v1

import (
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func getTodayStartAndEndTime() (startTime string, endTime string) {
	startTime = time.Now().Format("2006-01-02 00:00")
	endTime = (time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Now().Location())).Format("2006-01-02 15:04:05")
	return
}

func GetPerformance(context *gin.Context) {
	startTime, endTime := getTodayStartAndEndTime()
	StackResponse, err := services.GetStackPerformance(startTime, endTime)
	QuotaResponse, err := services.GetQuotaData(startTime, endTime)
	PagePerformanceListResponse, err := services.GetLoadInfoPageList(startTime, endTime)
	StageTimeResponse, err := services.GetStageTimeList(startTime, endTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.PagePerformanceResponse{
			QuotaResponse:               QuotaResponse,
			StackResponse:               StackResponse,
			PagePerformanceListResponse: PagePerformanceListResponse,
			StageTimeResponse:           StageTimeResponse,
		}, "获取成功", context)
	}
}
