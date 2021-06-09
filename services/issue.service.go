package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"danci-api/utils"
)

func FindJsIssue(message string) (jsIssues model.Issue, err error) {
	err = global.GVA_DB.Where("message = ? ", message).Find(&jsIssues).Error
	return
}

func GetIssues() (pageJsErrorList []response.PageJsErrorList, err error) {
	var issueList []response.PageJsErrorList
	err = global.GVA_DB.Model(&model.Issue{}).Select("issues.id, " +
		"issues.error_name, " +
		"issues.message, " +
		"COUNT(DISTINCT page_issues.user_id) as error_user, " +
		"COUNT(page_issues.id) as total").Joins("" +
		"INNER JOIN page_issues on page_issues.issues_id = issues.id" +
		"").Group("issues.id").Find(&issueList).Error

	startTime, endTime := utils.GetTodayStartAndEndTime();
	for _, issue := range issueList {
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("page_url, created_at as first_time").Where("issues_id = ? ", &issue.ID).Group("id ASC").Limit(1).Scan(&issue).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("created_at as last_time").Where("issues_id = ? ", &issue.ID).Group("id DESC").Limit(1).Scan(&issue).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("COUNT(id) as today").Where("from_unixtime(page_issues.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&issue).Error
		pageJsErrorList = append(pageJsErrorList, issue)
	}
	return
}

func GetIssuesDetail(issueId, errorId int, monitorId string) (jsErrorDetail response.PageJsErrorDetail, err error) {
	if errorId != 0 {
		err = global.GVA_DB.Model(&model.PageIssue{}).Where("id = ? And monitor_id = ?", errorId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as previous_error_id ").Where("id < ? And monitor_id = ? ", errorId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail.PreviousErrorID).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? ", errorId, monitorId).Group("id").Limit(1).Scan(&jsErrorDetail.NextErrorID).Error
	} else {
		err = global.GVA_DB.Model(&model.PageIssue{}).Where("issues_id = ?", issueId).Group("id DESC").Limit(1).Scan(&jsErrorDetail).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as previous_error_id").Where("id < ? And monitor_id = ? ", jsErrorDetail.ID, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail.PreviousErrorID).Error
		err = global.GVA_DB.Model(&model.PageIssue{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? ", jsErrorDetail.ID, monitorId).Group("id").Limit(1).Scan(&jsErrorDetail.NextErrorID).Error
	}
	return
}
