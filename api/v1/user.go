package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	responses, err := services.GetUsers()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetUser(context *gin.Context) {
	var userRequest request.UserRequest
	_ = context.BindQuery(&userRequest)
	responses, err := services.GetUser(userRequest.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetUserActions(context *gin.Context) {
	responses, err := services.GetUserActions()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetUserAction(context *gin.Context) {
	var userActionRequest request.UserActionRequest
	_ = context.BindQuery(&userActionRequest)
	var responses interface{}
	var err error
	if userActionRequest.ActionType == "PAGE_LOAD" {
		responses, err = services.GetActionPerformance(userActionRequest.ActionID)
	} else if userActionRequest.ActionType == "HTTP_LOG" {
		responses, err = services.GetActionHttp(userActionRequest.ActionID)
	} else if userActionRequest.ActionType == "JS_ERROR" {
		responses, err = services.GetActionJsError(userActionRequest.ActionID)
	} else if userActionRequest.ActionType == "RESOURCE_ERROR" {
		responses, err = services.GetActionResourceError(userActionRequest.ActionID)
	}
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
