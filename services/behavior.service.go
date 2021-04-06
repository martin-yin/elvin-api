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

func GetBehaviors() (behaviorsResponse []response.BehaviorsResponse, err error) {
	err = global.GVA_DB.Model(&model.UserAction{}).Find(&behaviorsResponse).Error
	return
}

func GetBehaviorHttp(id string) (behavior response.BehaviorHttpResponse, err error) {
	err = global.GVA_DB.Model(&model.PageHttp{}).Where("id = ?", id).Scan(&behavior).Error
	return
}

func GetBehaviorPerformance(id string) (behavior response.BehaviorPerformanceResponse, err error) {
	err = global.GVA_DB.Model(&model.PagePerformance{}).Where("id = ?", id).Scan(&behavior).Error
	return
}

func GetBehaviorError(id string) (userInfo UserInfo, err error) {
	err = global.GVA_DB.Model(&model.PagePerformance{}).Where("id = ?", id).Scan(&userInfo).Error
	return
}

//func GetBehaviorPerformance() (behavior response.UserInfo, err error){
//	err = global.GVA_DB.Model(&model.UserBehaviorInfo{}).Find(&behavior).Error
//	return
//}
