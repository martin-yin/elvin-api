package model

import "danci-api/global"

type PageHttp struct {
	global.GVA_MODEL
	PageUrl      string      `json:"page_url"`
	HttpUrl      string      `json:"http_url" gorm:"index:http_url"`
	LoadTime     float64     `json:"load_time"`
	Method       string      `json:"method"`
	Status       int         `json:"status"`
	StatusText   string      `json:"status_text"`
	StatusResult string      `json:"status_result"`
	RequestText  string      `json:"request_text"  gorm:"type:text"`
	ResponseText string      `json:"response_text" gorm:"type:text"`
	PublicFiles  PublicFiles `gorm:"embedded"`
}
