package response

type SurveyStatisticsResponse struct {
	Pv        int    `json:"pv"`
	Uv        int    `json:"uv"`
	Ip        string `json:"ip"`
	Resources int    `json:"resources"`
	JsError   int    `json:"js_error"`
}

type LoadPageList struct {
	LoadPage int `json:"load_page"`
	TimeKey  int `json:"time_key"`
}

type SurveyPerformancesResponse struct {
	LoadPage float64 `json:"load_page"`
	TimeKey  string  `json:"time_key"`
}
