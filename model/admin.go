package model

import "danci-api/global"

type Admin struct {
	global.GVA_MODEL
	UserName string `json:"user_name"`
	Password string `json:"password" gorm:"-"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
	//Teams []Team `gorm:"many2many:team_admins;"`
}
