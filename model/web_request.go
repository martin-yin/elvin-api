package model

import (
	"danci-api/global"
)

// 页面请求
type WebRequest struct {
	global.GVA_MODEL
	GenerateTime int    `json:"generate_time" gorm:"上报时间"`
	Method       string `json:"method" gorm:"comment: 请求方式"`
	HttpType     string `json:"http_type"`
	ElapsedTime  int    `json:"elapsed_time"`
	Code         int    `json:"code" gorm:"comment: 响应状态码"`
	ApiUrl       string `json:"api_url" gorm:"comment: api_url"`
	PageUrl      string `json:"page_url gorm:"comment: 页面url"`
	Message      string `json:"message"`
	IsError      int    `json:"is_error"`
	Body         string `json:"body" gorm:"comment: 请求体, 如果是post请求的话有"`
	// 运营商 网络 UA 浏览器类型 操作系统 地区 和 IP地址

	// 浏览器类型 操作系统 可以通过UserAgent 解析出来。
	UserAgent string `json:"user_agent"`
	IPAddress string `json:"ip_address"`
}

// PV 计算当前这一分钟有多少个用户
