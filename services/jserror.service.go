package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetJsError() (pageJsErrorList []response.PageJsErrorList, err error) {
	err = global.GVA_DB.Model(&model.PageJsError{}).Select("id, COUNT(DISTINCT id) as frequency, stack, message").Group("message").Scan(&pageJsErrorList).Error
	return
}

func GetJsErrorDetail(id string) (jsErrorDetail model.PageJsError, err error) {
	err = global.GVA_DB.Model(&model.PageJsError{}).Where("id = ?", id).First(&jsErrorDetail).Error
	return
}
