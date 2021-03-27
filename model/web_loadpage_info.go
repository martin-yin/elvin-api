package model

import "danci-api/global"

// 页面load 信息记录
type WebLoadpageInfo struct {
	global.GVA_MODEL

	LoadPage     int `json:"load_page"`
	DomReady     int `json:"dom_ready"`
	Redirect     int `json:"redirect"`
	LookupDomain int `json:"lookup_domain"`
	Ttfb         int `json:"ttfb"`
	Request      int `json:"request"`
	LoadEvent    int `json:"load_event"`
	Appcache     int `json:"appcache"`
	UnloadEvent  int `json:"unload_event"`
	Connect      int `json:"connect"`
	LoadType     int `json:"load_type"`
	BrowserInfo  int `json:"browser_info"`
}
