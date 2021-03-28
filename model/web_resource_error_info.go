package model

import "danci-api/global"

// 页面资源错误
type WebResourceErrorInfo struct {
	global.GVA_MODEL

	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`
	HappenDate float64 `json:"happen_date"`
	PageKey    string  `json:"page_key"`

	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`
}
