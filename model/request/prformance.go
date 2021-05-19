package request

import "danci-api/model"

type PerformanceBody struct {
	PageUrl      string  `json:"page_url"`
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

	model.PublicFiles
}

type PerformanceParams struct {
	TimeGrain string `form:"time_grain"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	StageType string `form:"stage_type"`
	MonitorId string `form:"monitor_id"`
}
