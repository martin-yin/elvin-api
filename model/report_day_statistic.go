package model

// 每日数据记录
type ReportDayStatistic struct {
	MODEL
	ActionType string `json:"action_type"`
	MonitorId  string `json:"monitor_id"`
	Day        string `json:"day"`
	Count      string `json:"count"`
}
