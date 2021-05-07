package model

import "danci-api/global"

// 用户行为记录
type User struct {
	global.GVA_MODEL
	PublicFiles PublicFiles `gorm:"embedded"`
	UserId      string      `json:"user_id" gorm:"unique"`
}
