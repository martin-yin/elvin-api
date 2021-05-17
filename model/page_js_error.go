package model

import "danci-api/global"

// 页面js错误
//type PageJsError struct {
//	global.GVA_MODEL
//
//	ErrorName string `json:"error_name"`
//	Message   string `json:"message" gorm:"uniqueIndex"`
//	JsErrorStacks []JsErrorStack `json:"js_error_stacks" gorm:"foreignKey:PageJsErrorId"`
//}

//type PageJsError struct {
//	global.GVA_MODEL
//
//	ErrorName string `json:"error_name"`
//	Message   string `json:"message" gorm:"index:unique"`
//	JsErrorStacks []JsErrorStack `gorm:"foreignKey:PageJsErrorId"`
//}
//
//type JsErrorStack struct {
//	global.GVA_MODEL
//	// 对应归属到哪个PageJsError
//	PageJsErrorId uint
//	ComponentName string      `json:"componentName"`
//	Stack         string      `json:"stack" gorm:"type:text"`
//	ErrorName     string      `json:"error_name"`
//	PublicFiles   PublicFiles `gorm:"embedded"`
//}

type PageJsError struct {
	global.GVA_MODEL
	PageUrl       string      `json:"page_url"`
	ComponentName string      `json:"componentName"`
	Message       string      `json:"message"`
	Stack         string      `json:"stack" gorm:"type:text"`
	ErrorName     string      `json:"error_name"`
	PublicFiles   PublicFiles `json:"public_files" gorm:"embedded"`
	JsIssuesId    uint        `json:"js_issues_id"`
}

type JsIssue struct {
	global.GVA_MODEL
	ErrorName   string        `json:"error_name"`
	Message     string        `json:"message" gorm:"uniqueIndex"`
	PageJsError []PageJsError `json:"page_js_error" gorm:"foreignKey:JsIssuesId"`
}
