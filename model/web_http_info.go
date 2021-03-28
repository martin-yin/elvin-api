package model

import "danci-api/global"

type WebHttpInfo struct {
	global.GVA_MODEL

	PageUrl       string `json:"page_url"`
	SimpleHttpUrl string `json:"simple_http_url"`
	HttpUrl       string `json:"http_url"`

	UserId     string `json:"user_id"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`
	HappenDate string `json:"happen_date"`
	PageKey    string `json:"page_key"`

	LoadTime     string `json:"load_time"`
	Status       string `json:"status"`
	StatusText   string `json:"status_text"`
	StatusResult string `json:"status_result"`
	RequestText  string `json:"request_text"`
	ResponseText string `json:"response_text"`
}
