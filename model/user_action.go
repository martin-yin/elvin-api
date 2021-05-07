package model

import "danci-api/global"

// 用户行为记录
type UserAction struct {
	global.GVA_MODEL
	UserId       string `json:"user_id"`
	MonitorId    string `json:"monitor_id"`
	HappenTime   int    `json:"happen_time"`
	HappenDay    string `json:"happen_day"`
	ActionType   string `json:"action_type"`
	ActionID     uint   `json:"action_id"`
	ActionDetail string `json:"action_detail" gorm:"type:text"`
}
