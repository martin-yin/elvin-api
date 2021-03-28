package model

import "danci-api/global"

type Admin struct {
	global.GVA_MODEL

	UserName string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	NickName string `json:"nick_name"`
}
