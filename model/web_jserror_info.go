package model

import "danci-api/global"

// 页面js错误
type WebJsErrorInfo struct {
	global.GVA_MODEL

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"brower_name"`
	BrowserVersion string `json:"brower_version"`
	MonitorIp      string `json:"monitorIp"`
	InfoType       string `json:"info_type"`
	ErrorMessage   string `json:"error_message"`
	ErrorStack     string `json:"error_stack"`
	BrowserInfo    string `json:"browser_info"`
}
