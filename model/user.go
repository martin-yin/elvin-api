package model

import "danci-api/global"

// 用户行为记录
type User struct {
	global.GVA_MODEL
	UserId      string      `json:"user_id" gorm:"unique"`
	PublicFiles PublicFiles `gorm:"embedded"`
}
