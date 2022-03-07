package request

import "dancin-api/model"

type UserActionRequest struct {
	ActionID   string `form:"action_id"`
	ActionType string `form:"action_type"`
	StartTime  string `form:"start_time"`
	EndTime    string `form:"end_time"`
}

type UserRequest struct {
	ID string `form:"id"`
}

type UserActionsRequest struct {
	SessionId string `form:"session_id"`
	Page      int    `form:"page"`
	Limit     int    `form:"limit"`
}

type OperationBody struct {
	PageUrl     string `json:"page_url"`
	ClassName   string `json:"class_name"`
	Placeholder string `json:"placeholder"`
	InputValue  string `json:"input_value"`
	TagName     string `json:"tag_name"`
	InnerText   string `json:"inner_text"`
	Path        string `json:"path"`
	model.CommonFiles
}
