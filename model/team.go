package model

import "danci-api/global"

type Team struct {
	global.GVA_MODEL
	Name     string `json:"name"`
	LeaderId uint   `json:"leader_id"`

	Admins []Admin `gorm:"many2many:team_admins;"`
	Projects []Project `json:"team_projects"`
	//Admins   []Admin   `gorm:"many2many:team_admins;"`
	//Projects []Project `gorm:"many2many:team_projects;"`
}
