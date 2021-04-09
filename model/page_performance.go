package model

import "danci-api/global"

// 页面load 信息记录
type PagePerformance struct {
	global.GVA_MODEL

	PageUrl      string      `json:"page_url"`
	Redirect     float64     `json:"redirect"`
	Appcache     float64     `json:"appcache"`
	LookupDomain float64     `json:"lookup_domain"`
	Tcp          float64     `json:"tcp"`
	SslT         float64     `json:"ssl_t"`
	Request      float64     `json:"request"`
	DomParse     float64     `json:"dom_parse"`
	Ttfb         float64     `json:"ttfb"`
	LoadPage     float64     `json:"load_page"`
	LoadEvent    float64     `json:"load_event"`
	LoadType     string      `json:"load_type"`
	PublicFiles  PublicFiles `gorm:"embedded"`
}
