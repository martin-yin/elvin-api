package model

import "danci-api/global"

// 页面load 信息记录
type WebLoadpageInfo struct {
	global.GVA_MODEL

	PageUrl string `json:"page_url"`
	UserId      string    `json:"user_id"`
	UploadType  string    `json:"upload_type"`
	HappenTime  string    `json:"happen_time"`
	HappenDate  string    `json:"happen_date"`
	PageKey     string `json:"page_key"`

	LoadPage     string `json:"load_page"`
	DomReady     int64 `json:"dom_ready"`
	Redirect     int64 `json:"redirect"`
	LookupDomain int64 `json:"lookup_domain"`
	Ttfb         int64 `json:"ttfb"`
	Request     float64 `json:"Request"`
	LoadEvent    int64 `json:"load_event"`
	Appcache     int64 `json:"appcache"`
	UnloadEvent  float64 `json:"unload_event"`
	Connect      int64 `json:"connect"`
	LoadType     string `json:"load_type"`
	BrowserInfo  string `json:"browser_info"`
}
