package model

import "danci-api/global"

// 用户行为记录
type User struct {
	global.GVA_MODEL
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	HappenTime     int    `json:"happen_time" gorm:"unique"`
	HappenDay      string `json:"happen_day"`
	EventId        string `json:"event_id"`
	IP             string `json:"ip"`
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
