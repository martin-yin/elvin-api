package model

import (
	"danci-api/global"
)

// 页面请求
type WebResourcesError struct {
	global.GVA_MODEL

	PageUrl      string `json:"page_url"`
	ResourceUrl  string `json:"resource_url"`
	GenerateTime int    `json:"generate_time"`
	DomPath      string `json:"dom_path"`

	// 浏览器类型 操作系统 可以通过UserAgent 解析出来。
	UserAgent string `json:"user_agent"`
	IPAddress string `json:"ip_address"`
}
