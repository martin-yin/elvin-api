package request

import "danci-api/model"

type PageViewBody struct {
	PageUrl        string `json:"page_url"`

	model.PublicFiles
}
