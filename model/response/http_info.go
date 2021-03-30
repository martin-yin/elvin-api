package response

type HttpInfoListResponse struct {
	PageUrl      string  `json:"page_url"`
	RequestUrl   string  `json:"request_url"`
	LoadTime     float64 `json:"load_time"`
	SuccessRate  string  `json:"success_rate"`
	RequestTotal int     `json:"request_total"`
}

type HttpQuotaResponse struct {
	RequestTotal int    `json:"request_total"`
	LoadTime     int    `json:"load_time"`
	SuccessRate  string `json:"success_rate"`
	ErrorUser    int    `json:"error_user"`
}

type WebHttpInfoResponse struct {
	HttpInfoListResponse []HttpInfoListResponse `json:"http_info_list"`
	HttpQuotaResponse    HttpQuotaResponse      `json:"http_quota"`
}
