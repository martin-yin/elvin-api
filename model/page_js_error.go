package model

import "danci-api/global"

// 页面js错误
type PageJsError struct {
	global.GVA_MODEL

	ErrorName string `json:"error_name"`
	Message   string `json:"message" gorm:"uniqueIndex"`
	JsErrorStacks []JsErrorStack `gorm:"foreignKey:PageJsErrorId"`
}

type JsErrorStack struct {
	global.GVA_MODEL
	// 对应归属到哪个PageJsError
	PageJsErrorId uint
	ComponentName string      `json:"componentName"`
	Stack         string      `json:"stack" gorm:"type:text"`
	ErrorName     string      `json:"error_name"`
	PublicFiles   PublicFiles `gorm:"embedded"`
}
