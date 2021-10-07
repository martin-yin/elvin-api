package services

import (
	"dancin-api/global"
	"dancin-api/model"
)

func GetUserDetail() (userInfo []interface{}, err error) {
	err = global.GORMDB.Model(&model.UserAction{}).Find(&userInfo).Error
	return
}
