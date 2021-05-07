package v1

import (
	"danci-api/global"
	"danci-api/model/request"
	"danci-api/model/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func CreatePagePerformance(context *gin.Context) {
	var performanceBody request.PerformanceBody
	err := context.ShouldBindJSON(&performanceBody)
	performanceBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(performanceBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := performanceBody.HappenDay + performanceBody.MonitorId + performanceBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreateHttpInfo(context *gin.Context) {
	var httpBody request.HttpBody
	err := context.BindJSON(&httpBody)
	httpBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(httpBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := httpBody.HappenDay + httpBody.MonitorId
	if httpBody.Status >= 400 {
		incrKey = "HTTP_ERROR_LOG"
	} else {
		incrKey = "HTTP_LOG"
	}
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreateResourcesError(context *gin.Context) {
	var resourceErroBody request.ResourceErrorBody
	err := context.BindJSON(&resourceErroBody)
	resourceErroBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(resourceErroBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := resourceErroBody.HappenDay + resourceErroBody.MonitorId + resourceErroBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreatePageOperation(context *gin.Context) {
	var operationBody request.OperationBody
	err := context.BindJSON(&operationBody)
	operationBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(operationBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := operationBody.HappenDay + operationBody.MonitorId + operationBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreatePageJsError(context *gin.Context) {
	var jsErrorBody request.JsErrorBody
	err := context.BindJSON(&jsErrorBody)
	jsErrorBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(jsErrorBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := jsErrorBody.HappenDay + jsErrorBody.MonitorId + jsErrorBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreatePageView(context *gin.Context) {
	var pageViewBody request.PageViewBody
	err := context.BindJSON(&pageViewBody)
	pageViewBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pageViewBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := pageViewBody.HappenDay + pageViewBody.MonitorId + pageViewBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}
