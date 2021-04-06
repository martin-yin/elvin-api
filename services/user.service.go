package services

import (
	"danci-api/global"
	"danci-api/model"
)

func GetUserDetail() (userInfo []interface{}, err error) {
	err = global.GVA_DB.Model(&model.UserAction{}).Find(&userInfo).Error
	return
}
