package model

import "danci-api/global"

type WebHttpInfo struct {
	global.GVA_MODEL

	// 页面url
	PageUrl       string `json:"page_url"`
	// http请求url
	HttpUrl       string `json:"http_url"`

	// 用户id
	UserId     string `json:"user_id"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`
	HappenDate string `json:"happen_date"`
	PageKey    string `json:"page_key"`

	LoadTime     float64 `json:"load_time"`
	Status       int `json:"status"`
	StatusText   string `json:"status_text"`
	StatusResult string `json:"status_result"`
	RequestText  string `json:"request_text"`
	ResponseText string `json:"response_text"`
}
