package model

import "danci-api/global"

type WebHttpInfo struct {
	global.GVA_MODEL

	// 页面url
	PageUrl string `json:"page_url"`
	// http请求url
	HttpUrl string `json:"http_url" gorm:"index:http_url"`

	// 用户id
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`

	LoadTime     float64 `json:"load_time"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"  gorm:"type:text"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type WebHttpInfoStatistical struct {
	global.GVA_MODEL
	// 页面url
	PageUrl string `json:"page_url"`
	// http请求url
	HttpUrl string `json:"http_url" gorm:"index:http_url"`

	SuccessTotal int `json:"success_total"`

	FailTotal int `json:"fail_total"`

	Total int `json:"fail_total"`
}
