package response

type UserResponse struct {
	ID             string `json:"id"`
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	HappenTime     string `json:"happen_time" gorm:"unique"`
	IP             string `json:"ip" gorm:"unique"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
}

type BehaviorPerformanceResponse struct {
	PageUrl    string  `json:"page_url"`
	Request    float64 `json:"request"`
	DomParse   float64 `json:"dom_parse"`
	Ttfb       float64 `json:"ttfb"`
	LoadPage   float64 `json:"load_page"`
	LoadEvent  float64 `json:"load_event"`
	LoadType   string  `json:"load_type"`
	UploadType string  `json:"upload_type"`
	HappenTime string  `json:"happen_time"`
}

type BehaviorHttpResponse struct {
	PageUrl      string  `json:"page_url"`
	UserId       string  `json:"user_id"`
	ApiKey       string  `json:"api_key"`
	UploadType   string  `json:"upload_type"`
	HappenTime   string  `json:"happen_time"`
	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"`
}

type BehaviorsResponse struct {
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	HappenTime string `json:"happen_time"`
	// 打开页面 发送请求 点击等等……
	BehaviorType string `json:"behavior_type"`
	BehaviorId   uint   `json:"behavior_id"`
	// 如果是打开页面的话
	PageUrl  string `json:"page_url"`
	LoadType string `json:"load_type"`
	// 如果是 http请求的话
	HttpUrl string `json:"http_url"`
	// 如果是资源加载失败的话
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	// 如果是发生错误的话
	Message string `json:"message"`
	Stack   string `json:"stack"`
	//如果是点击
	ClassName  string `json:"class_name"`
	InnterText string `json:"innter_text" gorm:"type:text"`
}
