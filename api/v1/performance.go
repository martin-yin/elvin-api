package v1

import (
	"danci-api/model/request"
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
	var queryPagePerformance request.QueryPagePerformance
	err := context.BindQuery(&queryPagePerformance)
	StackResponse, err := services.GetStackPerformance(queryPagePerformance.StartTime, queryPagePerformance.EndTime)
	QuotaResponse, err := services.GetQuotaData(queryPagePerformance.StartTime, queryPagePerformance.EndTime)
	PagePerformanceListResponse, err := services.GetLoadInfoPageList(queryPagePerformance.StartTime, queryPagePerformance.EndTime)
	StageTimeResponse, err := services.GetStageTimeList(queryPagePerformance.StartTime, queryPagePerformance.EndTime, queryPagePerformance.TimeGrain)
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
