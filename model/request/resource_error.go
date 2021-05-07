package request

type ResourceErrorBody struct {
	PageUrl     string `json:"page_url"`
	UserId      string `json:"user_id"`
	MonitorId   string `json:"monitor_id"`
	ActionType  string `json:"action_type"`
	HappenDay   string `json:"happen_day"`
	HappenTime  int    `json:"happen_time"`
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`
	EventId     string `json:"event_id"`
	// 设备信息
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type ResourceErrorParams struct {
	TimeGrain string `form:"time_grain"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	StageType string `form:"stage_type"`
	MonitorId string `form:"monitor_id"`
}
