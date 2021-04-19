package request

type PostJsErrorBody struct {
	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	ActionType string `json:"action_type"`
	HappenDay  string `json:"happen_time"`
	HappenTime int    `json:"happen_day"`
	// 组件名称
	ComponentName  string `json:"component_name"`
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
