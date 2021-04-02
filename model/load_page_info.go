package model

import "danci-api/global"

// 页面load 信息记录
type LoadpageInfo struct {
	global.GVA_MODEL
	PageUrl      string  `json:"page_url"`
	UserId       string  `json:"user_id"`
	ApiKey       string  `json:"api_key"`
	UploadType   string  `json:"upload_type"`
	HappenTime   string  `json:"happen_time"`
	Redirect     float64 `json:"redirect"`
	Appcache     float64 `json:"appcache"`
	LookupDomain float64 `json:"lookup_domain"`
	Tcp          float64 `json:"tcp"`
	SslT         float64 `json:"ssl_t"`
	Request      float64 `json:"request"`
	DomParse     float64 `json:"dom_parse"`
	Ttfb         float64 `json:"ttfb"`
	LoadPage     float64 `json:"load_page"`
	LoadEvent    float64 `json:"load_event"`
	LoadType     string  `json:"load_type"`

	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"brower_version"`
	UA             string `json:"ua"`
}
