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
	incrKey := pagePerformanceBody.HappenDay + pagePerformanceBody.MonitorId + pagePerformanceBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func Report(context *gin.Context) {
	//actionType := context.PostForm("action_type")
	//if actionType == "PAGE_LOAD" {
	//	CreatePagePerformance(context)
	//}
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
	incrKey := pageHttpBody.HappenDay + pageHttpBody.MonitorId
	if pageHttpBody.Status >= 400 {
		incrKey = "HTTP_ERROR_LOG"
	} else {
		incrKey = "HTTP_LOG"
	}
	global.GVA_REDIS.Incr(incrKey)
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
	incrKey := pageResourceErroBody.HappenDay + pageResourceErroBody.MonitorId + pageResourceErroBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
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
	incrKey := behaviorBody.HappenDay + behaviorBody.MonitorId + behaviorBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
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
	incrKey := jsErrorBody.HappenDay + jsErrorBody.MonitorId + jsErrorBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
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
	incrKey := pageViewBody.HappenDay + pageViewBody.MonitorId + pageViewBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}
