package model

type UserAction struct {
	MODEL
	UserId       string `json:"user_id"`
	MonitorId    string `json:"monitor_id"`
	HappenTime   int    `json:"happen_time"`
	HappenDay    string `json:"happen_day"`
	ActionType   string `json:"action_type"`
	SessionId    string `json:"session_id"`
	ActionDetail string `json:"action_detail" gorm:"type:text"`
}
