package request

import "dancin-api/model"

type PageViewBody struct {
	PageUrl       string `json:"page_url"`
	DocumentTitle string `json:"document_title"`
	Referrer      string `json:"referrer"`
	Encode        string `json:"encode"`
	model.CommonFiles
}
