package response

type UserResponse struct {
	ID             string `json:"id"`
	UserId         string `json:"user_id"`
	MonitorId      string `json:"monitor_id"`
	HappenTime     string `json:"happen_time" gorm:"unique"`
	IP             string `json:"ip" gorm:"unique"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	SessionId      string `json:"session_id"`
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
	MonitorId      string  `json:"monitor_id"`
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
	MonitorId  string `json:"monitor_id"`
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
	MonitorId      string `json:"monitor_id"`
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
	MonitorId      string `json:"monitor_id"`
	HappenTime     int    `json:"happen_time"`
	HappenDay      string `json:"happen_day"`
	ActionType     string `json:"action_type"`
	ClassName      string `json:"class_name"`
	Placeholder    string `json:"placeholder"`
	InputValue     string `json:"Input_value"`
	TagName        string `json:"tag_name"`
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
	MonitorId      string `json:"monitor_id"`
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

type ActionsResponse struct {
	UserId       string `json:"user_id"`
	MonitorId    string `json:"monitor_id"`
	HappenTime   int    `json:"happen_time"`
	HappenDay    string `json:"happen_day"`
	ActionType   string `json:"action_type"`
	SessionId    string `json:"session_id"`
	ActionDetail string `json:"action_detail" gorm:"type:text"`
}

type ActionsStatisticsResponse struct {
	ActionType string `json:"action_type"`
	Total      string `json:"total"`
}

type UserActionsResponse struct {
	ActionsResponse []ActionsResponse `json:"user_actions_list"`
	Total           int               `json:"total"`
	Page            int               `json:"page"`
}
