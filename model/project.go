package model

import "danci-api/global"

type Project struct {
	global.GVA_MODEL

	ProjectName string `json:"project"`
	AppKey      string `json:"app_key"`

	AdminId string `json:"admin_id"`
	Email   string `json:"email"`
}
