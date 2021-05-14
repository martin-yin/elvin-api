package request

type JsErrorBody struct {
	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	MonitorId  string `json:"monitor_id"`
	ActionType string `json:"action_type"`
	HappenTime  int    `json:"happen_time"`
	HappenDay   string `json:"happen_day"`
	ErrorName  string `json:"error_name"`
	// 组件名称
	ComponentName  string `json:"component_name"`
	Line           string `json:"line"`
	Column         string `json:"column"`
	Stack          string `json:"stack"`
	Message        string `json:"message"`
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

type JsErrorParams struct {
	ID string `form:"id"`
}
