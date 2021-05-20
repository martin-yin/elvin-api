package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
)

func FindJsIssue(message string) (jsIssues model.Issue, err error) {
	err = global.GVA_DB.Where("message = ? ", message).Find(&jsIssues).Error
	return
}

func GetIssues() (pageJsErrorList []response.PageJsErrorList, err error) {
	err = global.GVA_DB.Model(&model.Issue{}).Select("js_issues.id, " +
		"js_issues.error_name, " +
		"js_issues.message, " +
		"COUNT(DISTINCT page_js_errors.user_id) as error_user, " +
		"COUNT(page_js_errors.id) as error_count").Joins("" +
		"INNER JOIN page_js_errors on page_js_errors.js_issues_id = js_issues.id" +
		"").Find(&pageJsErrorList).
		Group("js_issues.id").Error
	return
}

func GetIssuesDetail(issueId, errorId int, monitorId string) (jsErrorDetail response.PageJsErrorDetail, err error) {
	if errorId != 0 {
		err = global.GVA_DB.Model(&model.PageIssue{}).Where("id = ? And monitor_id = ?", errorId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as previous_error_id ").Where("id < ? And monitor_id = ? ", errorId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail.PreviousErrorID).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? ", errorId, monitorId).Group("id").Limit(1).Scan(&jsErrorDetail.NextErrorID).Error
	} else {
		err = global.GVA_DB.Model(&model.PageIssue{}).Where("js_issues_id = ?", issueId).Group("id DESC").Limit(1).Scan(&jsErrorDetail).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as previous_error_id").Where("id < ? And monitor_id = ? ", jsErrorDetail.ID, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail.PreviousErrorID).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? ", jsErrorDetail.ID, monitorId).Group("id").Limit(1).Scan(&jsErrorDetail.NextErrorID).Error
	}
	return
}
