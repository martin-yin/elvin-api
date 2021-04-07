package model

import "danci-api/global"

// 页面js错误
type PageJsError struct {
	global.GVA_MODEL

	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	ActionType string `json:"action_type"`
	HappenTime int    `json:"happen_time"`
	HappenDay  string `json:"happen_day"`
	// 组件名称
	ComponentName string `json:"componentName"`
	Stack         string `json:"stack" gorm:"type:text"`
	Message       string `json:"message"`

	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
