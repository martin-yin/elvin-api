package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetUserBehaviors(context *gin.Context)  {
	responses, err := services.GetBehaviors()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetUserBehavior(context *gin.Context) {
	var behaviorRequest request.BehaviorRequest
	_ = context.BindQuery(&behaviorRequest)
	var responses interface{}
	var err error
	if behaviorRequest.BehaviorType == "PAGE_LOAD" {
		responses, err = services.GetBehaviorPerformance(behaviorRequest.BehaviorId)
	} else if  behaviorRequest.BehaviorType == "HTTP_LOG" {
		responses, err = services.GetBehaviorHttp(behaviorRequest.BehaviorId)
	}
	//responses, err := services.GetPerformance(behaviorRequest.BehaviorId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}