package response

type ProjectResponse struct {
	ID          string `json:"id"`
	ApiKey      string `json:"api_key"`
	AdminId     string `json:"admin_id"`
	ProjectName string `json:"project_name"`
	Logo        string `json:"logo"`
}

type HomeStatisticsDataProjectResponse struct {
	ID                string `json:"id"`
	ApiKey            string `json:"api_key"`
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
