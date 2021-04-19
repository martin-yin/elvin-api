package model

import "danci-api/global"

type PageHttp struct {
	global.GVA_MODEL
	PageUrl      string      `json:"page_url"`
	HttpUrl      string      `json:"http_url" gorm:"index:http_url"`
	LoadTime     float64     `json:"load_time"`
	Status       int         `json:"status"`
	StatusText   string      `json:"status_text"`
	StatusResult string      `json:"status_result"`
	RequestText  string      `json:"request_text"`
	ResponseText string      `json:"response_text" gorm:"type:text"`
	PublicFiles  PublicFiles `gorm:"embedded"`
}

type PageHttpStatistical struct {
	global.GVA_MODEL
	PageUrl      string `json:"page_url"`
	HttpUrl      string `json:"http_url" gorm:"index:http_url"`
	SuccessTotal int    `json:"success_total"`
	FailTotal    int    `json:"fail_total"`
	Total        int    `json:"fail_total"`
	HappenDay    string `json:"happen_day"`
}
