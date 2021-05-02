package model

import "danci-api/global"

// 每日数据记录
type ReportDayStatistic struct {
	global.GVA_MODEL
	ActionType string `json:"action_type"`
	MonitorId  string `json:"monitor_id"`
	Day        string `json:"day"`
	Count      string    `json:"count"`
}
