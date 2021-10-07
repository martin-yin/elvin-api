package request

import "dancin-api/model"

type PageViewBody struct {
	PageUrl string `json:"page_url"`

	model.CommonFiles
}
