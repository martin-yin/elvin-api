package request

import "danci-api/model"

type OperationBody struct {
	PageUrl     string `json:"page_url"`
	ClassName   string `json:"class_name"`
	Placeholder string `json:"placeholder"`
	InputValue  string `json:"Input_value"`
	TagName     string `json:"tag_name"`
	InnerText   string `json:"inner_text"`

	model.PublicFiles
}
