package response

type HttpListResponse struct {
	PageUrl      string  `json:"page_url"`
	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Total        int     `json:"total"`
	FailTotal    int     `json:"fail_total"`
	SuccessTotal int     `json:"success_total"`
}

type HttpQuotaResponse struct {
	Total        int `json:"total"`
	LoadTime     int `json:"load_time"`
	SuccessTotal int `json:"success_total"`
	ErrorUser    int `json:"error_user"`
}

type PageHttpResponse struct {
	HttpListResponse  []HttpListResponse `json:"http_info_list"`
	HttpQuotaResponse HttpQuotaResponse  `json:"http_quota"`
}
