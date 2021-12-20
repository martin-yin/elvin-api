package response

type HttpListResponse struct {
	Url             string  `json:"url"`
	LoadTime        float64 `json:"load_time"`
	Total           int     `json:"total"`
	Status          int     `json:"status"`
	UserTotal       int     `json:"user_total"`
	ErrorUser       int     `json:"error_user"`
	UserSlow        int     `json:"user_slow"`
	SuccessRate     float64 `json:"success_rate"`
	LastHappenTime  int     `json:"last_happen_time"`
	FirstHappenTime int     `json:"first_happen_time"`
}

type HttpErrsResponse struct {
	Url             string `json:"url"`
	Total           int    `json:"total"`
	Status          int    `json:"status"`
	UserTotal       int    `json:"user_total"`
	LastHappenTime  int    `json:"last_happen_time"`
	FirstHappenTime int    `json:"first_happen_time"`
	RequestText     string `json:"request_text"`
}

type HttpQuotaResponse struct {
	Total    float64 `json:"total"`
	LoadTime float64 `json:"load_time"`
	// 请求成功的条数
	SuccessTotal float64 `json:"success_total"`
	// 请求的用户个数
	ErrorUser int `json:"error_user"`
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
	HttpListResponse  []HttpListResponse `json:"http_list"`
	HttpQuotaResponse HttpQuotaResponse  `json:"http_quota"`
}

type PageHttpStage struct {
	HttpStageTimeResponse []HttpStageTimeResponse `json:"http_stagetime"`
}

type PageHttpStageError struct {
	HttpStageTimeResponseError []HttpStageTimeResponseError `json:"http_stagetime"`
}

type HttpErrorQuotaResponse struct {
	Error400  float64 `json:"error_400"`
	Error404  int     `json:"error_404"`
	Error500  int     `json:"error_500"`
	ErrorUser int     `json:"error_user"`
}

type PageHttpErrorResponse struct {
	HttpListResponse       []HttpListResponse     `json:"http_error_list"`
	HttpErrorQuotaResponse HttpErrorQuotaResponse `json:"http_error_quota"`
}
