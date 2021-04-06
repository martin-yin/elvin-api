package model

import "danci-api/global"

// 页面js错误
type PageJsError struct {
	global.GVA_MODEL

	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`
	// 组件名称
	ComponentName  string `json:"componentName"`
	Stack          string `json:"stack" gorm:"type:text"`
	Message        string `json:"message"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
