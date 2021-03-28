package model

import "danci-api/global"

// 页面资源错误
type WebResourceErrorInfo struct {
	global.GVA_MODEL


	PageUrl string `json:"page_url"`
	UserId      string    `json:"user_id"`
	UploadType  string    `json:"upload_type"`
	HappenTime  string    `json:"happen_time"`
	HappenDate  string    `json:"happen_date"`
	PageKey     string `json:"page_key"`

	SourceUrl   int    `json:"source_url"`
	ElementType int    `json:"element_type"`
	Status      int    `json:"status"`
}
