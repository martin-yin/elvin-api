package model

import "danci-api/global"

// 用户行为记录
type UserBehaviorInfo struct {
	global.GVA_MODEL

	UserId     string `json:"user_id"`
	ApiKey     string `json:"api_key"`
	HappenTime string `json:"happen_time"`
	// 打开页面 发送请求 点击等等……
	BehaviorType string `json:"behavior_type"`

	BehaviorId uint `json:"behavior_id"`
	// 如果是打开页面的话
	PageUrl  string `json:"page_url"`
	LoadType string `json:"load_type"`

	// 如果是 http请求的话
	HttpUrl string `json:"http_url"`

	// 如果是资源加载失败的话
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`

	// 如果是发生错误的话
	Message string `json:"message"`
	Stack string `json:"stack"`

	//如果是点击
	ClassName    string `json:"class_name"`
	InnterText   string `json:"innter_text" gorm:"type:text"`
}
