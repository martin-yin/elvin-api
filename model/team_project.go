package model

import "danci-api/global"

type TeamProject struct {
	global.GVA_MODEL
	Member  uint `json:"member"`
	Project uint `json:"project"`
}
