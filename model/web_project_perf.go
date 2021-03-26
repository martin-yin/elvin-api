package model

import (
	"danci-api/global"
)

// 页面性能与项目关联表
type WebProjectPerf struct {
	global.GVA_MODEL
	ProjectId int `json:"project_id" gorm:"comment: 项目ID"`
	PerfId    int `json:"perf_id" gorm:"comment: perf_id"`
}
