package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserList(context *gin.Context) {
	var userRequest request.UsersRequest
	err := context.BindQuery(&userRequest)
	responses, err := services.GetUsers(userRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetUser(context *gin.Context) {
	var userRequest request.UserRequest
	_ = context.BindQuery(&userRequest)
	responses, err := services.GetUser(userRequest.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetUsersActionsStatistics(context *gin.Context) {
	var userActionsRequest request.UserActionsRequest
	_ = context.BindQuery(&userActionsRequest)
	actionStatisticsResponse, err := services.GetUserActionsStatistics(userActionsRequest.SessionId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(actionStatisticsResponse, "获取成功", context)
}

func GetUserActionList(context *gin.Context) {
	var userActionsRequest request.UserActionsRequest
	_ = context.BindQuery(&userActionsRequest)
	fmt.Print("userActionsRequest", userActionsRequest.SessionId)
	actionResponse, err := services.GetUserActions(userActionsRequest.SessionId, userActionsRequest.Page, userActionsRequest.Limit)
	total, err := services.GetUserActionsTotal(userActionsRequest.SessionId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(response.UserActionsResponse{
		ActionsResponse: actionResponse,
		Total:           total,
		Page:            userActionsRequest.Page,
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
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
