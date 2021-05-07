package request

type TeamParam struct {
	Name string `json:"name"`
}

type BindTeamAdminsParams struct {
	TeamId   uint   `form:"team_id" json:"team_id"`
	AdminIds string `form:"admin_ids" json:"admin_ids"`
}

type AddTeamProjectParams struct {
	TeamId    uint `form:"team_id" json:"team_id"`
	ProjectId uint `form:"project_id" json:"project_id"`
	ProjectName string `form:"project_name" json:"project_name"`
	ProjectType string `form:"project_type" json:"project_type"`
	Logo        string `form:"logo" json:"logo"`
}

type ProjectParams struct {
	ProjectName string `form:"project_name" json:"project_name"`
	ProjectType string `form:"project_type" json:"project_type"`
	Logo        string `form:"logo" json:"logo"`
	MonitorId   string `form:"monitor_id" json:"monitor_id"`
	TeamId      uint   `form:"team_id" json:"team_id"`
}

type AddTeamParams struct {
	Name string `json:"name"`
}
