package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"fmt"
)

func Login(loginParam request.Login) (err error, userInter *model.Admin) {
	var user model.Admin
	err = global.GORMDB.Model(&model.Admin{}).Where("user_name = ? AND password = ?", loginParam.Username, loginParam.Password).First(&user).Error
	return err, &user
}

func RegisterAdmin(admin model.Admin) (userInter *model.Admin, err error) {
	err = global.GORMDB.Model(&model.Admin{}).Create(&admin).Error
	return &admin, err
}

func FindAdmins(adminIds ...uint) (adminList []model.Admin, err error) {
	err = global.GORMDB.Model(&model.Admin{}).Where("id in ?", adminIds).Find(&adminList).Error
	fmt.Println(adminList, "adminList")
	return
}
