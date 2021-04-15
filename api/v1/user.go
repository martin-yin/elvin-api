package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	var userRequest request.UsersRequest
	err := context.BindQuery(&userRequest)
	responses, err := services.GetUsers(userRequest)
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
	var userActionsRequest request.UserActionsRequest
	_ = context.BindQuery(&userActionsRequest)
	actionResponse, err := services.GetUserActions(userActionsRequest.EventID)
	actionStatisticsResponse, err := services.GetUserActionsStatistics(userActionsRequest.EventID)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(response.UserActionsResponse{
		BehaviorsResponse:           actionResponse,
		BehaviorsStatisticsResponse: actionStatisticsResponse,
	}, "获取成功", context)
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
	} else if userActionRequest.ActionType == "BEHAVIOR_INFO" {
		responses, err = services.GetActionBehavior(userActionRequest.ActionID)
	} else if userActionRequest.ActionType == "PAGE_VIEW" {
		responses, err = services.GetActionPageView(userActionRequest.ActionID)
	}
	if err != nil {
		response.FailWithMessage(err.Error(), context)
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
