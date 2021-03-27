package model

import "danci-api/global"

// page得点击记录
type WebBehaviorInfo struct {
	global.GVA_MODEL

	BehaviorType int    `json:"behavior_type"`
	ClassName    int    `json:"class_name"`
	Placeholder  int    `json:"placeholder"`
	InputValue   int    `json:"Input_value"`
	TagNameint   string `json:"tag_name"`
	InnerText    int    `json:"inner_text"`
}
