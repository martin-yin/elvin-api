package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/utils"
)

type WhereSql struct {
	sql string
}

func GetUsers(params request.RequestParams) (userResponse []response.UserResponse, err error) {
	sql, paramSql := utils.BuildWhereSql("users", "", params)
	err = global.GORMDB.Model(&model.User{}).Where(sql, paramSql...).Group("happen_time desc").Find(&userResponse).Error
	return
}

func GetUser(id string) (userResponse response.UserResponse, err error) {
	err = global.GORMDB.Model(&model.User{}).Where("id = ?", id).Find(&userResponse).Error
	return
}

func GetUserActions(sessionId string, page int, limit int) (actionResponse []response.ActionsResponse, err error) {
	err = global.GORMDB.Model(&model.UserAction{}).Where("session_id = ?", sessionId).Order("happen_time").Limit(limit).Offset((page - 1) * limit).Find(&actionResponse).Error
	return
}

func GetUserActionsTotal(sessionId string) (total int, err error) {
	err = global.GORMDB.Model(&model.UserAction{}).Select("count(*) as total").Where("session_id = ?", sessionId).Order("happen_time").Find(&total).Error
	return
}

func GetUserActionsStatistics(sessionId string) (actionStatisticsResponse []response.ActionsStatisticsResponse, err error) {
	err = global.GORMDB.Model(&model.UserAction{}).Select("action_type, count(*) as total").Where("session_id = ?", sessionId).Group("action_type").Find(&actionStatisticsResponse).Error
	return
}

func GetActionHttp(id string) (actionResponse response.ActionHttpResponse, err error) {
	err = global.GORMDB.Model(&model.PageHttp{}).Where("id = ?", id).Scan(&actionResponse).Error
	return
}

func GetActionPerformance(id string) (actionPerformanceResponse response.ActionPerformanceResponse, err error) {
	err = global.GORMDB.Model(&model.PagePerformance{}).Where("id = ?", id).Scan(&actionPerformanceResponse).Error
	return
}

func GetActionJsError(id string) (actionJsErrorResponse response.ActionJsErrorResponse, err error) {
	err = global.GORMDB.Model(&model.PageIssue{}).Where("id = ?", id).Scan(&actionJsErrorResponse).Error
	return
}

func GetActionResourceError(id string) (actionResourceErrorResponse response.ActionResourceErrorResponse, err error) {
	err = global.GORMDB.Model(&model.PageResourceError{}).Where("id = ?", id).Scan(&actionResourceErrorResponse).Error
	return
}

func GetActionBehavior(id string) (actionPageBehaviorResponse response.ActionPageBehaviorResponse, err error) {
	err = global.GORMDB.Model(&model.PageOperation{}).Where("id = ?", id).Scan(&actionPageBehaviorResponse).Error
	return
}

func GetActionPageView(id string) (actionPageViewResponse response.ActionPageViewResponse, err error) {
	err = global.GORMDB.Model(&model.PageView{}).Where("id = ?", id).Scan(&actionPageViewResponse).Error
	return
}
