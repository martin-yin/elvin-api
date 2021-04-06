package model

import "danci-api/global"

// 用户行为记录
type User struct {
	global.GVA_MODEL
	UserId string `json:"user_id"`
	ApiKey string `json:"api_key"`

	HappenTime     string `json:"happen_time" gorm:"unique"`
	IP             string `json:"ip" gorm:"unique"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
}
