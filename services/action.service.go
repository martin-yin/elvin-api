package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetUsers() (userResponse []response.UserResponse, err error) {
	err = global.GVA_DB.Model(&model.User{}).Find(&userResponse).Error
	return
}

func GetUser(id string) (userResponse response.UserResponse, err error) {
	err = global.GVA_DB.Model(&model.User{}).Where("id = ?", id).Find(&userResponse).Error
	return
}

func GetUserActions() (actionResponse []response.BehaviorsResponse, err error) {
	err = global.GVA_DB.Model(&model.UserAction{}).Find(&actionResponse).Error
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

//
//func GetBehaviorError(id string) (userInfo UserInfo, err error) {
//	err = global.GVA_DB.Model(&model.PagePerformance{}).Where("id = ?", id).Scan(&userInfo).Error
//	return
//}

//func GetBehaviorPerformance() (behavior response.UserInfo, err error){
//	err = global.GVA_DB.Model(&model.UserBehaviorInfo{}).Find(&behavior).Error
//	return
//}
