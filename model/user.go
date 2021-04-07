package model

import "danci-api/global"

// 用户行为记录
type User struct {
	global.GVA_MODEL
	UserId         string `json:"user_id"`
	ApiKey         string `json:"api_key"`
	HappenTime     int    `json:"happen_time" gorm:"unique"`
	HappenDay      string `json:"happen_day" gorm:"unique"`
	IP             string `json:"ip" gorm:"unique"`
	Device         string `json:"device" gorm:"unique"`
	DeviceType     string `json:"device_type" gorm:"unique"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
