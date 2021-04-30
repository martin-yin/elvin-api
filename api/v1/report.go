package v1

import (
	"danci-api/global"
	"danci-api/model/request"
	"danci-api/model/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func CreatePagePerformance(context *gin.Context) {
	var pagePerformanceBody request.PostPagePerformance
	err := context.ShouldBindJSON(&pagePerformanceBody)
	pagePerformanceBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pagePerformanceBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

func CreateHttpInfo(context *gin.Context) {
	var pageHttpBody request.PostPageHttpBody
	err := context.BindJSON(&pageHttpBody)
	pageHttpBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pageHttpBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

func CreateResourcesError(context *gin.Context) {
	var pageResourceErroBody request.PostPageResourceErroBody
	err := context.BindJSON(&pageResourceErroBody)
	pageResourceErroBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pageResourceErroBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

func CreatePageBehavior(context *gin.Context) {
	var behaviorBody request.PostBehaviorInfoBody
	err := context.BindJSON(&behaviorBody)
	behaviorBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(behaviorBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

func CreatePageJsError(context *gin.Context) {
	var jsErrorBody request.PostJsErrorBody
	err := context.BindJSON(&jsErrorBody)
	jsErrorBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(jsErrorBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

func CreatePageView(context *gin.Context) {
	var pageViewBody request.PostPageViewBody
	err := context.BindJSON(&pageViewBody)
	pageViewBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pageViewBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}
