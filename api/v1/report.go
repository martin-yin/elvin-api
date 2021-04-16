package v1

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func CreatePagePerformance(context *gin.Context) {
	var pagePerformanceBody request.PostPagePerformance
	err := context.ShouldBindJSON(&pagePerformanceBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pagePerformanceBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

// 存储HTTP请求
func CreateHttpInfo(context *gin.Context) {
	var pageHttpBody request.PostPageHttpBody
	err := context.BindJSON(&pageHttpBody)
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
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(behaviorBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}

func CreatePageJsError(context *gin.Context) {
	var jsErrorInfoBody request.PostJsErrorInfoBody
	var publicFiles model.PublicFiles
	files, _ := context.Get("public_files")
	err := utils.StructToJsonToStruct(files, &publicFiles)
	err = context.BindJSON(&jsErrorInfoBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	jsErrorInfoModel := model.PageJsError{
		PageUrl:       jsErrorInfoBody.PageUrl,
		ComponentName: jsErrorInfoBody.ComponentName,
		Stack:         jsErrorInfoBody.Stack,
		Message:       jsErrorInfoBody.Message,
		PublicFiles:   publicFiles,
	}
	if err := services.CreatePageJsError(jsErrorInfoModel, jsErrorInfoBody.EventId); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

func CreatePageView(context *gin.Context) {
	var pageViewBody request.PostPageViewBody
	err := context.BindJSON(&pageViewBody)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pageViewBody)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
}
