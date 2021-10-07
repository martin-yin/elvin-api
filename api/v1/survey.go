package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"dancin-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetSurveyStatisticsData(context *gin.Context) {
	var surveyParams request.SurveyParams
	err := context.BindQuery(&surveyParams)
	startTime, endTime := utils.GetTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyStatisticsData(startTime, endTime, surveyParams.MonitorId.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}

func GetSurveyPUvData(context *gin.Context) {
	var surveyParams request.SurveyParams
	err := context.BindQuery(&surveyParams)
	startTime, endTime := utils.GetTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyPUvData(startTime, endTime, surveyParams.MonitorId.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}

func GetSurveyJsErrorData(context *gin.Context) {
	var surveyParams request.SurveyParams
	err := context.BindQuery(&surveyParams)
	startTime, endTime := utils.GetTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyJsErrorData(startTime, endTime, surveyParams.MonitorId.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}
