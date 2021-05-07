package request

type HttpBody struct {
	PageUrl      string  `json:"page_url"`
	UserId       string  `json:"user_id"`
	MonitorId    string  `json:"monitor_id"`
	ActionType   string  `json:"action_type"`
	HappenTime   int     `json:"happen_time"`
	HappenDay    string  `json:"happen_day"`
	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"`
	// 设备信息
	EventId        string `json:"event_id"`
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type HttpParams struct {
	TimeGrain string `form:"time_grain"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	StageType string `form:"stage_type"`
	MonitorId string `form:"monitor_id"`
}
