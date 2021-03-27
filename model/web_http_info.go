package model

import "danci-api/global"

type WebHttpInfo struct {
	global.GVA_MODEL

	HttpUrl       string `json:"http_url"`
	SimpleHttpUrl string `json:"simple_http_url"`
	LoadTime      string `json:"load_time"`
	Status        string `json:"status"`
	StatusText    string `json:"status_text"`
	StatusResult  string `json:"status_result"`
	RequestText   string `json:"request_text"`
	ResponseText  string `json:"response_text"`
}
