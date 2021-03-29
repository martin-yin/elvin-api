package model

import "danci-api/global"

// 页面load 信息记录
type WebLoadpageInfo struct {
	global.GVA_MODEL

	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`
	PageKey    string  `json:"page_key"`

	DomReady     float64 `json:"dom_ready"`
	Redirect     float64 `json:"redirect"`
	LookupDomain float64 `json:"lookup_domain"`
	Ttfb         float64 `json:"ttfb"`
	Request      float64 `json:"Request"`
	LoadEvent    float64 `json:"load_event"`
	Appcache     float64 `json:"appcache"`
	UnloadEvent  float64 `json:"unload_event"`
	Connect      float64 `json:"connect"`
	LoadType     string  `json:"load_type"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
