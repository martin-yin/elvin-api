package model

import (
	"danci-api/global"
)

// 页面性能
type WebPerformance struct {
	global.GVA_MODEL
	Appcache        float64 `json:"appcache" `
	Contentdownload float64 `json:"contentdownload"`
	Dns             float64 `json:"dns"`
	Domparsing      float64 `json:"domparsing"`
	PageUrl         string  `json:"page_url"`
	Redirect        float64 `json:"redirect"`
	Res             float64 `json:"res"`
	Tcp             float64 `json:"tcp" `
	Ttfb            float64 `json:"ttfb"`

	Cooike    string `json:"cookie"`
	UserAgent string `json:"user_agent"`
	IPAddress string `json:"ip_address"`
	//Information
}

// PV 计算当前这一分钟有多少个用户
