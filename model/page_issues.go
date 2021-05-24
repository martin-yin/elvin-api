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

type PageIssue struct {
	global.GVA_MODEL
	PageUrl       string      `json:"page_url"`
	ComponentName string      `json:"componentName"`
	Message       string      `json:"message"`
	Stack         string      `json:"stack" gorm:"type:text"`
	StackFrames   string      `json:"stack_frames" gorm:"type:text"`
	ErrorName     string      `json:"error_name"`
	PublicFiles   PublicFiles `json:"public_files" gorm:"embedded"`
	IssuesId      uint        `json:"issues_id"`
}

type Issue struct {
	global.GVA_MODEL
	ErrorName  string      `json:"error_name"`
	Message    string      `json:"message" gorm:"uniqueIndex"`
	HappenTime int         `json:"happen_time"`
	MonitorId  string      `json:"monitor_id"`
	PageIssue  []PageIssue `json:"page_issue" gorm:"foreignKey:IssuesId"`
}
