package model

import "danci-api/global"

// page得点击记录
type WebBehaviorInfo struct {
	global.GVA_MODEL

	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`
	HappenDate string `json:"happen_date"`
	PageKey    string `json:"page_key"`

	BehaviorType int    `json:"behavior_type"`
	ClassName    int    `json:"class_name"`
	Placeholder  int    `json:"placeholder"`
	InputValue   int    `json:"Input_value"`
	TagNameint   string `json:"tag_name"`
	InnerText    int    `json:"inner_text"`
}
