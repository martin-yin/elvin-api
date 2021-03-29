package request

type WebLoadPageInfo struct {
	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`

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

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type WebHttpInfo struct {
	PageUrl    string  `json:"url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`

	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type WebResourceErrorInfo struct {
	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`

	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type WebBehaviorInfo struct {
	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`

	BehaviorType string `json:"behavior_type"`
	ClassName    string `json:"class_name"`
	Placeholder  string `json:"placeholder"`
	InputValue   string `json:"Input_value"`
	TagNameint   string `json:"tag_name"`
	InnterText   string `json:"innter_text"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
