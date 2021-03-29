package model

import "danci-api/global"

// page得点击记录
type WebBehaviorInfo struct {
	global.GVA_MODEL

	PageUrl    string  `json:"page_url"`
	UserId     string  `json:"user_id"`
	UploadType string  `json:"upload_type"`
	HappenTime float64 `json:"happen_time"`
	HappenDate float64 `json:"happen_date"`
	PageKey    string  `json:"page_key"`

	BehaviorType string `json:"behavior_type"`
	ClassName    string `json:"class_name"`
	Placeholder  string `json:"placeholder"`
	InputValue   string `json:"Input_value"`
	TagNameint   string `json:"tag_name"`
	InnterText   string `json:"innter_text" gorm:"type:text"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}
