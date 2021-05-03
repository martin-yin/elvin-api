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
	fmt.Println(err, "error!!!!!!!!!!!!!!!!!!!")
	return err, &user
}
