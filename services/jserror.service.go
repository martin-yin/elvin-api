package services

import (
	"danci-api/global"
	"danci-api/model"
	"gorm.io/gorm"
)

func GetJsError() (pageJsErrorList []model.PageJsError, err error) {
	err = global.GVA_DB.Preload("JsErrorStacks", func(db *gorm.DB) *gorm.DB {
		return db.Select("js_error_stacks.stack")
	}).Model(&model.PageJsError{}).Find(&pageJsErrorList).Error
	return
}

func GetJsErrorDetail(id string) (jsErrorDetail model.PageJsError, err error) {
	err = global.GVA_DB.Model(&model.PageJsError{}).Where("id = ?", id).First(&jsErrorDetail).Error
	return
}
