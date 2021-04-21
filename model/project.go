package model

import "danci-api/global"

type Project struct {
	global.GVA_MODEL
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	Logo        string `json:"logo"`
	TeamId      string `json:"team_id"`
	AdminId     string `json:"admin_id"`
	MonitorId string `json:"monitor_id"`
}
