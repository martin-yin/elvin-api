package model

import "danci-api/global"

type Project struct {
	global.GVA_MODEL
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	Logo        string `json:"logo"`
	MonitorId   string `json:"monitor_id"`
	AdminID     uint   `json:"admin_id"`
	TeamID      uint   `json:"team_id"`
}
