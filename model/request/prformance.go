package request

import "danci-api/model"

type ReportBody struct {
	ActionType string `json:"action_type" form: "action_type"`
	HappenTime int    `json:"happen_time" form: "happen_time"`
	HappenDay  string `json:"happen_day" form: "happen_day"`
	MonitorId  string `json:"monitor_id" form: "monitor_id"`
	IP         string `json:"ip"`
}

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
	StageType string `form:"stage_type"`
	MonitorId
	StartEndTime
}
