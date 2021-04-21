package model

import "danci-api/global"

// 用户行为记录
type UserAction struct {
	global.GVA_MODEL
	UserId     string `json:"user_id"`
	MonitorId  string `json:"monitor_id"`
	HappenTime int    `json:"happen_time"`
	HappenDay  string `json:"happen_day"`
	// 打开页面 发送请求 点击等等……
	ActionType string `json:"action_type"`
	ActionID   uint   `json:"action_id"`
	// 如果是打开页面的话
	PageUrl  string `json:"page_url"`
	LoadType string `json:"load_type"`
	EventId  string `json:"event_id"`
	// 如果是 http请求的话
	HttpUrl string `json:"http_url"`
	// 如果是资源加载失败的话
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	// 如果是发生错误的话
	Message string `json:"message"`
	Stack   string `json:"stack"`
	//如果是点击
	ClassName  string `json:"class_name"`
	InnterText string `json:"innter_text" gorm:"type:text"`
}
