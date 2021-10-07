package model

type Project struct {
	MODEL
	ProjectName string `json:"project_name" gorm:"unique"`
	ProjectType string `json:"project_type"`
	Logo        string `json:"logo"`
	MonitorId   string `json:"monitor_id" gorm:"unique"`
	AdminID     uint   `json:"admin_id"`
	TeamID      uint   `json:"team_id"`
}
