package model

import "danci-api/global"

// 页面资源错误
type PageResourceError struct {
	global.GVA_MODEL
	PageUrl     string      `json:"page_url"`
	SourceUrl   string      `json:"source_url"`
	ElementType string      `json:"element_type"`
	Status      string      `json:"status"`
	PublicFiles PublicFiles `gorm:"embedded"`
}
