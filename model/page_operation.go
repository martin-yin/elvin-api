package model

import "danci-api/global"

// page得点击记录
type PageOperation struct {
	global.GVA_MODEL
	PageUrl     string      `json:"page_url"`
	ClassName   string      `json:"class_name"`
	Placeholder string      `json:"placeholder"`
	InputValue  string      `json:"Input_value"`
	TagName     string      `json:"tag_name"`
	InnerText   string      `json:"inner_text" gorm:"type:text"`
	PublicFiles PublicFiles `gorm:"embedded"`
}
