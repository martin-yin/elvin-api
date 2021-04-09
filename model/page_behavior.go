package model

import "danci-api/global"

// page得点击记录
type PageBehavior struct {
	global.GVA_MODEL
	PageUrl     string      `json:"page_url"`
	ClassName   string      `json:"class_name"`
	Placeholder string      `json:"placeholder"`
	InputValue  string      `json:"Input_value"`
	TagNameint  string      `json:"tag_name"`
	InnterText  string      `json:"innter_text" gorm:"type:text"`
	PublicFiles PublicFiles `gorm:"embedded"`
}
