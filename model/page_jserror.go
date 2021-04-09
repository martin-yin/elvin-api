package model

import "danci-api/global"

// 页面js错误
type PageJsError struct {
	global.GVA_MODEL
	PageUrl string `json:"page_url"`
	// 组件名称
	ComponentName string      `json:"componentName"`
	Stack         string      `json:"stack" gorm:"type:text"`
	Message       string      `json:"message"`
	PublicFiles   PublicFiles `gorm:"embedded"`
}
