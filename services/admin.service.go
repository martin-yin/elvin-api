package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/request"
	"fmt"
)

func Login(loginParam request.Login) (err error, userInter *model.Admin) {
	var user model.Admin
	err = global.GVA_DB.Model(&model.Admin{}).Where("user_name = ? AND password = ?", loginParam.Username, loginParam.Password).First(&user).Error
	return err, &user
}

func RegisterAdmin(admin model.Admin) (userInter *model.Admin, err error) {
	err = global.GVA_DB.Model(&model.Admin{}).Create(&admin).Error
	return &admin, err
}

func FindAdmins(adminIds ...uint) (adminList []model.Admin, err error) {
	err = global.GVA_DB.Model(&model.Admin{}).Where("id in ?", adminIds).Find(&adminList).Error
	fmt.Println(adminList, "adminList")
	return
}
