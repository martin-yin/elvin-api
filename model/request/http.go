package request

import "dancin-api/model"

type HttpBody struct {
	PageUrl      string  `json:"page_url"`
	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Method       string  `json:"method"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"`

	model.CommonFiles
}

type HttpParams struct {
	TimeGrain string `form:"time_grain"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	StageType string `form:"stage_type"`
	MonitorId string `form:"monitor_id"`
}
