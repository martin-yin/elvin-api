package model

import "danci-api/global"

// 页面js错误
type JsErrorInfo struct {
	global.GVA_MODEL

	PageUrl      string `json:"page_url"`
	UserId       string `json:"user_id"`
	ApiKey       string `json:"api_key"`
	UploadType   string `json:"upload_type"`
	HappenTime   string `json:"happen_time"`
	MonitorIp    string `json:"monitorIp"`
	InfoType     string `json:"info_type"`
	ErrorMessage string `json:"error_message"`
	ErrorStack   string `json:"error_stack"`

	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"brower_version"`
	UA             string `json:"ua"`
}
