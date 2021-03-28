package request

type WebLoadPageInfo struct {
	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`
	HappenDate float64 `json:"happen_date"`
	PageKey    string  `json:"page_key"`

	DomReady     float64 `json:"dom_ready"`
	Redirect     float64 `json:"redirect"`
	LookupDomain float64 `json:"lookup_domain"`
	Ttfb         float64 `json:"ttfb"`
	Request      float64 `json:"request"`
	LoadEvent    float64 `json:"load_event"`
	Appcache     float64 `json:"appcache"`
	UnloadEvent  float64 `json:"unload_event"`
	Connect      float64 `json:"connect"`
	LoadType     string  `json:"load_type"`
	BrowserInfo  string  `json:"browser_info"`
}

type WebHttpInfo struct {
	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`
	HappenDate float64 `json:"happen_date"`
	PageKey    string  `json:"page_key"`

	HttpUrl       string `json:"http_url"`
	LoadTime      float64 `json:"load_time"`
	Status        int `json:"status"`
	StatusText    string `json:"status_text"`
	StatusResult  string `json:"status_result"`
	RequestText   string `json:"request_text"`
	ResponseText  string `json:"response_text"`
}

type WebResourceErrorInfo struct {
	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`
	HappenDate string `json:"happen_date"`
	PageKey    string `json:"page_key"`

	SourceUrl   int `json:"source_url"`
	ElementType int `json:"element_type"`
	Status      int `json:"status"`
}
