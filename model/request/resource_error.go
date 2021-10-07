package request

import "dancin-api/model"

type ResourceErrorBody struct {
	PageUrl     string `json:"page_url"`
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`

	model.CommonFiles
}

type ResourceErrorParams struct {
	TimeGrain string `form:"time_grain"`
	StageType string `form:"stage_type"`

	MonitorId
	StartEndTime
}
