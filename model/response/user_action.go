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
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	EventId        string `json:"event_id"`
	Nation         string `json:"nation"`
	Province       string `json:"province"`
	City           string `json:"city"`
	District       string `json:"district"`
}

type ActionPerformanceResponse struct {
	PageUrl        string  `json:"page_url"`
	Request        float64 `json:"request"`
	DomParse       float64 `json:"dom_parse"`
	Ttfb           float64 `json:"ttfb"`
	LoadPage       float64 `json:"load_page"`
	LoadEvent      float64 `json:"load_event"`
	LoadType       string  `json:"load_type"`
	ActionType     string  `json:"action_type"`
	HappenTime     string  `json:"happen_time"`
	Device         string  `json:"device"`
	DeviceType     string  `json:"device_type"`
	Os             string  `json:"os"`
	OsVersion      string  `json:"os_version"`
	Browser        string  `json:"browser"`
	BrowserVersion string  `json:"browser_version"`
	UA             string  `json:"ua"`
}

type ActionHttpResponse struct {
	PageUrl        string  `json:"page_url"`
	UserId         string  `json:"user_id"`
	ApiKey         string  `json:"api_key"`
	ActionType     string  `json:"action_type"`
	HappenTime     string  `json:"happen_time"`
	HttpUrl        string  `json:"http_url"`
	LoadTime       float64 `json:"load_time"`
	Status         int     `json:"status"`
	StatusText     string  `json:"status_text"`
	StatusResult   string  `json:"status_result"`
	RequestText    string  `json:"request_text"`
	ResponseText   string  `json:"response_text"`
	Device         string  `json:"device"`
	DeviceType     string  `json:"device_type"`
	Os             string  `json:"os"`
	OsVersion      string  `json:"os_version"`
	Browser        string  `json:"browser"`
	BrowserVersion string  `json:"browser_version"`
	UA             string  `json:"ua"`
}

type ActionJsErrorResponse struct {
	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	ActionType string `json:"action_type"`
	HappenTime string `json:"happen_time"`
	// 组件名称
	ComponentName  string `json:"componentName"`
	Stack          string `json:"stack" gorm:"type:text"`
	Message        string `json:"message"`
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type ActionResourceErrorResponse struct {
	PageUrl        string `json:"page_url"`
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	ActionType     string `json:"action_type"`
	HappenTime     string `json:"happen_time"`
	SourceUrl      string `json:"source_url"`
	ElementType    string `json:"element_type"`
	Status         string `json:"status"`
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type ActionPageBehaviorResponse struct {
	PageUrl        string `json:"page_url"`
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	HappenTime     int    `json:"happen_time"`
	HappenDay      string `json:"happen_day"`
	ActionType     string `json:"action_type"`
	ClassName      string `json:"class_name"`
	Placeholder    string `json:"placeholder"`
	InputValue     string `json:"Input_value"`
	TagNameint     string `json:"tag_name"`
	InnterText     string `json:"innter_text"`
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type ActionPageViewResponse struct {
	PageUrl        string `json:"page_url"`
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	HappenTime     int    `json:"happen_time"`
	HappenDay      string `json:"happen_day"`
	ActionType     string `json:"action_type"`
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type BehaviorsResponse struct {
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	HappenTime string `json:"happen_time"`
	// 打开页面 发送请求 点击等等……
	ActionType string `json:"action_type"`
	ActionID   uint   `json:"action_id"`
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

type BehaviorsStatisticsResponse struct {
	ActionType string `json:"action_type"`
	Total string `json:"total"`
}


type UserActionsResponse struct {
	BehaviorsResponse []BehaviorsResponse `json:"user_actions_list"`
	BehaviorsStatisticsResponse []BehaviorsStatisticsResponse `json:"user_action_statistics"`
}