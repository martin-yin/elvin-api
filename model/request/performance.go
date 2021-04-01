package request

type WebLoadPageInfo struct {
	PageUrl        string  `json:"page_url"`
	UserId         string  `json:"user_id"`
	ApiKey         string  `json:"api_key"`
	UploadType     string  `json:"upload_type"`
	HappenTime     string  `json:"happen_time"`
	Redirect       float64 `json:"redirect"`
	Appcache       float64 `json:"appcache"`
	LookupDomain   float64 `json:"lookup_domain"`
	Tcp            float64 `json:"tcp"`
	SslT           float64 `json:"ssl_t"`
	Request        float64 `json:"request"`
	DomParse       float64 `json:"dom_parse"`
	Ttfb           float64 `json:"ttfb"`
	LoadPage       float64 `json:"load_page"`
	LoadEvent      float64 `json:"load_event"`
	LoadType       string  `json:"load_type"`
	DeviceName     string  `json:"device_name"`
	Os             string  `json:"os"`
	BrowserName    string  `json:"browser_name"`
	BrowserVersion string  `json:"browser_version"`
	UA             string  `json:"ua"`
}

type WebHttpInfo struct {
	PageUrl        string  `json:"page_url"`
	UserId         string  `json:"user_id"`
	ApiKey         string  `json:"api_key"`
	UploadType     string  `json:"upload_type"`
	HappenTime     string  `json:"happen_time"`
	HttpUrl        string  `json:"http_url"`
	LoadTime       float64 `json:"load_time"`
	Status         int     `json:"status"`
	StatusText     string  `json:"status_text"`
	StatusResult   string  `json:"status_result"`
	RequestText    string  `json:"request_text"`
	ResponseText   string  `json:"response_text"`
	DeviceName     string  `json:"device_name"`
	Os             string  `json:"os"`
	BrowserName    string  `json:"browser_name"`
	BrowserVersion string  `json:"browser_version"`
	UA             string  `json:"ua"`
}

type WebResourceErrorInfo struct {
	PageUrl        string `json:"page_url"`
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	UploadType     string `json:"upload_type"`
	HappenTime     string `json:"happen_time"`
	SourceUrl      string `json:"source_url"`
	ElementType    string `json:"element_type"`
	Status         string `json:"status"`
	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type WebBehaviorInfo struct {
	PageUrl        string `json:"page_url"`
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	UploadType     string `json:"upload_type"`
	HappenTime     string `json:"happen_time"`
	BehaviorType   string `json:"behavior_type"`
	ClassName      string `json:"class_name"`
	Placeholder    string `json:"placeholder"`
	InputValue     string `json:"Input_value"`
	TagNameint     string `json:"tag_name"`
	InnterText     string `json:"innter_text"`
	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
