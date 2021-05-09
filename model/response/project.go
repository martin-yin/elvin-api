package response

type ProjectResponse struct {
	ID          string `json:"id"`
	MonitorId   string `form:"monitor_id" json:"monitor_id"`
	AdminId     string `json:"admin_id"`
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	Logo        string `json:"logo"`
}

type HomeStatisticsDataProjectResponse struct {
	ID                string `json:"id"`
	ApiKey            string `json:"api_key"`
	WebMonitorId      string `json:"web_monitor_id"`
	AdminId           string `json:"admin_id"`
	ProjectName       string `json:"project_name"`
	Logo              string `json:"logo"`
	NewUser           string `json:"new_user"`
	OldUser           string `json:"old_user"`
	TodayPv           string `json:"today_pv"`
	JsErrorRate       string `json:"js_error_rate"`
	HttpErrorRate     string `json:"http_error_rate"`
	ResourceErrorRate string `json:"resource_error_rate"`
}

type ProjectStatistics struct {
	Pv             float64 `json:"pv"`
	Uv             float64 `json:"uv"`
	JsError        float64 `json:"js_error"`
	ResourcesError float64 `json:"resources_error"`
	HttpError      float64 `json:"http_error"`
}
