package request

import "danci-api/model"

type IssuesBody struct {
	PageUrl       string `json:"page_url"`
	ErrorName     string `json:"error_name"`
	StackFrames   string `json:"stack_frames"`
	ComponentName string `json:"component_name"`
	Line          string `json:"line"`
	Column        string `json:"column"`
	Stack         string `json:"stack"`
	Message       string `json:"message"`

	model.PublicFiles
}

type JsErrorParams struct {
	IssueId   int    `form:"issue_id"`
	ErrorId   int    `form:"error_id"`
	MonitorId string `form:"monitor_id"`
}
