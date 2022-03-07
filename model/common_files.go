package model

import "time"

type MODEL struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommonFiles struct {
	UserId         string `json:"user_id"`
	MonitorId      string `json:"monitor_id"`
	ActionType     string `json:"action_type"`
	HappenTime     int    `json:"happen_time"`
	HappenDay      string `json:"happen_day"`
	IP             string `json:"ip"`
	SessionId      string `json:"session_id"`
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
	ConnectionType string `json:"connection_type"`
	Environment    string `json:"environment"`
	Language       string `json:"language"`
	Screen         string `json:"screen"`
	SdkVersion     string `json:"sdk_version"`
	Vp             string `json:"vp"`
}
