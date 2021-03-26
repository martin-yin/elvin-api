package model

import (
	"danci-api/global"
)

type WebProject struct {
	global.GVA_MODEL
	ApiKey      string `json:"u_ids" gorm:"comment: 用户列表"`
	ProjectName string
}
