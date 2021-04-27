package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetSurveyStatisticsData(context *gin.Context) {
	var querySuvey request.QueryPagePerformance
	err := context.BindQuery(&querySuvey)
	startTime, endTime := getTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyStatisticsData(startTime, endTime, querySuvey.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}

func GetSurveyPerformance(context *gin.Context)  {
	var querySuvey request.QueryPagePerformance
	err := context.BindQuery(&querySuvey)
	startTime, endTime := getTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyPerformance(startTime, endTime, querySuvey.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}