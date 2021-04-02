package model

import "danci-api/global"

// 页面资源错误
type ResourceErrorInfo struct {
	global.GVA_MODEL
	PageUrl     string `json:"page_url"`
	UserId      string `json:"user_id"`
	ApiKey      string `json:"api_key"`
	UploadType  string `json:"upload_type"`
	HappenTime  string `json:"happen_time"`
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`

	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"brower_version"`
	UA             string `json:"ua"`
}
