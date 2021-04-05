package services

import (
	"danci-api/global"
	"danci-api/model"
)

func GetUserDetail() (userInfo []interface{}, err error) {
	err = global.GVA_DB.Model(&model.UserBehaviorInfo{}).Find(&userInfo).Error
	return
}

