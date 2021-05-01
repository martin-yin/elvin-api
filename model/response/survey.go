package response

type SurveyStatisticsResponse struct {
	HttpError int     `json:"http_error"`
	LoadPage  float64 `json:"load_page"`
	Resources int     `json:"resources"`
	JsError   int     `json:"js_error"`
}

type SurveyPUvData struct {
	Pv      int `json:"pv"`
	UV      int `json:"uv"`
	TimeKey string `json:"time_key"`
}

type SurveyPerformancesResponse struct {
	LoadPage float64 `json:"load_page"`
	TimeKey  string  `json:"time_key"`
}
