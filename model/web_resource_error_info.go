package model

import "danci-api/global"

// 页面资源错误
type WebResourceErrorInfo struct {
	global.GVA_MODEL

	PageKey     string `json:"page_key"`
	SourceUrl   int    `json:"source_url"`
	ElementType int    `json:"element_type"`
	Status      int    `json:"status"`
}
