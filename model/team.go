package model

import "danci-api/global"

type Team struct {
	global.GVA_MODEL
	Name       string `json:"name"`
	LeaderId   string `json:"leader_id"`
	Members    string `json:"members"`
	MonitorIds string `json:"monitor_ids"`
}
