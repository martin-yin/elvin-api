package response

type HttpListResponse struct {
	PageUrl      string  `json:"page_url"`
	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Total        int     `json:"total"`
	UserTotal int 	`json:"user_total"`
}

type HttpQuotaResponse struct {
	Total        float64 `json:"total"`
	LoadTime     float64 `json:"load_time"`
}

type HttpStageTimeResponse struct {
	TimeKey     string  `json:"time_key"`
	Total       int     `json:"total"`
	SuccessRate float64 `json:"success_rate"`
	LoadTime    float64 `json:"load_time"`
}

type HttpStageTimeResponseError struct {
	TimeKey   string  `json:"time_key"`
	Total     int     `json:"total"`
	FailTotal float64 `json:"fail_total"`
	LoadTime  float64 `json:"load_time"`
}

type PageHttpResponse struct {
	HttpListResponse  []HttpListResponse `json:"http_info_list"`
	HttpQuotaResponse HttpQuotaResponse  `json:"http_quota"`
}

type PageHttpStage struct {
	HttpStageTimeResponse []HttpStageTimeResponse `json:"http_stagetime"`
}

type PageHttpStageError struct {
	HttpStageTimeResponseError []HttpStageTimeResponseError `json:"http_stagetime"`
}
