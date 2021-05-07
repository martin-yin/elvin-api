package request

type OperationBody struct {
	PageUrl     string `json:"page_url"`
	UserId      string `json:"user_id"`
	MonitorId   string `json:"monitor_id"`
	HappenTime  int    `json:"happen_time"`
	HappenDay   string `json:"happen_day"`
	ActionType  string `json:"action_type"`
	ClassName   string `json:"class_name"`
	Placeholder string `json:"placeholder"`
	InputValue  string `json:"Input_value"`
	TagNameint  string `json:"tag_name"`
	InnterText  string `json:"innter_text"`
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
