package model

import "danci-api/global"

type Team struct {
	global.GVA_MODEL
	Name string `json:"name"`

	NickName string `json:"nick_name"`
	AdminId  uint   `json:"admin_id"`

	Admins   []Admin   `gorm:"many2many:team_admins;" json:"team_admins"`
	Projects []Project `json:"team_projects"`
}
