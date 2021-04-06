package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func GetJsError() (pageJsErrorList []response.PageJsErrorList, err error) {
	err = global.GVA_DB.Model(&model.PageJsError{}).Select("COUNT(DISTINCT id) as frequency, message").Group("message").Scan(&pageJsErrorList).Error
	return
}
