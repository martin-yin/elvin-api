package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
)

type WhereSql struct {
	sql string
}

func GetUsers(usersParam request.UsersRequest) (userResponse []response.UserResponse, err error) {
	whereQuery := "from_unixtime(happen_time / 1000, '%Y-%m-%d %H:%i') between ? And ?"
	if usersParam.UserId != "" {
		whereQuery = whereQuery + " And user_id = ?"
	}
	startSearchTime := usersParam.SearchDate + " " + usersParam.SearchHour
	endSearchTime := usersParam.SearchDate + " 23:59:59"
	err = global.GVA_DB.Model(&model.User{}).Where(whereQuery, startSearchTime, endSearchTime, usersParam.UserId).Group("happen_time desc").Find(&userResponse).Error
	return
}

func GetUser(id string) (userResponse response.UserResponse, err error) {
	err = global.GVA_DB.Model(&model.User{}).Where("id = ?", id).Find(&userResponse).Error
	return
}

func GetUserActions(eventId string) (actionResponse []response.BehaviorsResponse, err error) {
	err = global.GVA_DB.Model(&model.UserAction{}).Where("event_id = ?", eventId).Order("happen_time").Find(&actionResponse).Error
	return
}

func GetActionHttp(id string) (actionResponse response.ActionHttpResponse, err error) {
	err = global.GVA_DB.Model(&model.PageHttp{}).Where("id = ?", id).Scan(&actionResponse).Error
	return
}

func GetActionPerformance(id string) (actionPerformanceResponse response.ActionPerformanceResponse, err error) {
	err = global.GVA_DB.Model(&model.PagePerformance{}).Where("id = ?", id).Scan(&actionPerformanceResponse).Error
	return
}

func GetActionJsError(id string) (actionJsErrorResponse response.ActionJsErrorResponse, err error) {
	err = global.GVA_DB.Model(&model.PageJsError{}).Where("id = ?", id).Scan(&actionJsErrorResponse).Error
	return
}

func GetActionResourceError(id string) (actionResourceErrorResponse response.ActionResourceErrorResponse, err error) {
	err = global.GVA_DB.Model(&model.PageResourceError{}).Where("id = ?", id).Scan(&actionResourceErrorResponse).Error
	return
}

func GetActionBehavior(id string) (actionPageBehaviorResponse response.ActionPageBehaviorResponse, err error) {
	err = global.GVA_DB.Model(&model.PageBehavior{}).Where("id = ?", id).Scan(&actionPageBehaviorResponse).Error
	return
}

func GetActionPageView(id string) (actionPageViewResponse response.ActionPageViewResponse, err error) {
	err = global.GVA_DB.Model(&model.PageView{}).Where("id = ?", id).Scan(&actionPageViewResponse).Error
	return
}

//
//func GetBehaviorError(id string) (userInfo UserInfo, err error) {
//	err = global.GVA_DB.Model(&model.PagePerformance{}).Where("id = ?", id).Scan(&userInfo).Error
//	return
//}

//func GetBehaviorPerformance() (behavior response.UserInfo, err error){
//	err = global.GVA_DB.Model(&model.UserBehaviorInfo{}).Find(&behavior).Error
//	return
//}
