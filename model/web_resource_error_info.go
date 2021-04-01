package model

import "danci-api/global"

// 页面资源错误
type WebResourceErrorInfo struct {
	global.GVA_MODEL

	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`

	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
