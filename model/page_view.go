package model

import "danci-api/global"

type PageView struct {
	global.GVA_MODEL
	PageUrl     string      `json:"page_url"`
	PublicFiles PublicFiles `gorm:"embedded"`
}
