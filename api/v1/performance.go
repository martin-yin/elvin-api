package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

// performance web性能相关的接口
func getPerformanceQuery(context *gin.Context) (performanceParams *request.PerformanceParams) {
	err := context.BindQuery(&performanceParams)
	performanceParams.StartTime = performanceParams.StartTime + " 00:00:00"
	performanceParams.EndTime = performanceParams.EndTime + " 23:59:59"
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	}
	return performanceParams
}

func GetPerformanceStack(context *gin.Context) {
	performanceParams := getPerformanceQuery(context)
	stack, err := services.GetPerformanceStack(performanceParams.MonitorId.MonitorId, performanceParams.StartTime, performanceParams.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(stack, "获取成功", context)
}

func GetPerformancePageList(context *gin.Context) {
	performanceParams := getPerformanceQuery(context)
	list, err := services.GetLoadInfoPageList(performanceParams.MonitorId.MonitorId, performanceParams.StartTime, performanceParams.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(list, "获取成功", context)

}

func GetPerformanceStageTime(context *gin.Context) {
	performanceParams := getPerformanceQuery(context)
	stageTime, err := services.GetStageTimeList(performanceParams.MonitorId.MonitorId, performanceParams.StartTime, performanceParams.EndTime, performanceParams.TimeGrain)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(stageTime, "获取成功", context)
}

func GetPerformanceQuota(context *gin.Context) {
	performanceParams := getPerformanceQuery(context)
	quota, err := services.GetQuotaData(performanceParams.MonitorId.MonitorId, performanceParams.StartTime, performanceParams.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(quota, "获取成功", context)
}

//

func GetPerformanceRankingList(context *gin.Context) {
	performanceParams := getPerformanceQuery(context)
	rankingList, err := services.GetRankingList(performanceParams.MonitorId.MonitorId, performanceParams.StartTime, performanceParams.EndTime)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
		return
	}
	response.OkWithDetailed(rankingList, "获取成功", context)
}
