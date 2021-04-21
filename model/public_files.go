package model

type PublicFiles struct {
	UserId         string `json:"user_id"`
	MonitorId   string `json:"monitor_id"`
	ActionType     string `json:"action_type"`
	HappenTime     int    `json:"happen_time"`
	HappenDay      string `json:"happen_day"`
	IP             string `json:"ip"`
	EventId        string `json:"event_id"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
	Nation         string `json:"nation"`
	Province       string `json:"province"`
	City           string `json:"city"`
	District       string `json:"district"`
}
