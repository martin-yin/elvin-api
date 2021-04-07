package model

import "danci-api/global"

// page得点击记录
type PageBehavior struct {
	global.GVA_MODEL

	PageUrl     string `json:"page_url"`
	UserId      string `json:"user_id"`
	ApiKey      string `json:"api_key"`
	HappenTime  int    `json:"happen_time"`
	HappenDay   string `json:"happen_day"`
	ActionType  string `json:"action_type"`
	ClassName   string `json:"class_name"`
	Placeholder string `json:"placeholder"`
	InputValue  string `json:"Input_value"`
	TagNameint  string `json:"tag_name"`
	InnterText  string `json:"innter_text" gorm:"type:text"`

	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}