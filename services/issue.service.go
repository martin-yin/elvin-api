package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/response"
	"fmt"
)

func FindJsIssue(message string) (jsIssues model.Issue, err error) {
	err = global.GORMDB.Where("message = ? ", message).Find(&jsIssues).Error
	return
}

func GetIssues(monitorId string, startTime string, endTime string) (pageJsErrorList []response.PageJsErrList, err error) {
	var issueList []response.PageJsErrList
	err = global.GORMDB.Model(&model.Issue{}).Select("issues.id, "+
		"issues.error_name, "+
		"issues.message, "+
		"page_issues.happen_time, page_issues.monitor_id, "+
		"COUNT(DISTINCT page_issues.user_id) as error_user, "+
		"COUNT(page_issues.id) as total").Joins(""+
		"INNER JOIN page_issues on page_issues.issues_id = issues.id"+
		"").Group("issues.id").Where(SqlWhereBuild("page_issues"), startTime, endTime, monitorId).Debug().Find(&issueList).Error

	fmt.Print("")

	for _, issue := range issueList {
		err = global.GORMDB.Model(&model.PageIssue{}).Select("page_url, created_at as first_time").Where("issues_id = ? ", &issue.ID).Group("id ASC").Limit(1).Scan(&issue).Error
		err = global.GORMDB.Model(&model.PageIssue{}).Select("created_at as last_time").Where("issues_id = ? ", &issue.ID).Group("id DESC").Limit(1).Scan(&issue).Error
		err = global.GORMDB.Model(&model.PageIssue{}).Select("COUNT(id) as today").Where("from_unixtime(page_issues.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime, endTime).Scan(&issue).Error
		pageJsErrorList = append(pageJsErrorList, issue)
	}
	return
}

func GetIssuesDetail(id, errId int, monitorId string) (jsErrDetail response.PageJsErrDetail, err error) {
	if errId != 0 && id == 0 {
		err = global.GORMDB.Model(&model.PageIssue{}).Where("issues_id = ? And monitor_id = ?", errId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrDetail).Error
		err = global.GORMDB.Model(&model.PageIssue{}).Select("id as previous_error_id ").Where("id < ? And monitor_id = ? And issues_id = ? ", jsErrDetail.ID, monitorId, errId).Group("id DESC").Limit(1).Scan(&jsErrDetail.PreviousErrorID).Error
		err = global.GORMDB.Model(&model.PageIssue{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? And issues_id = ? ", jsErrDetail.ID, monitorId, errId).Group("id").Limit(1).Scan(&jsErrDetail.NextErrorID).Error
	} else if id != 0 {
		err = global.GORMDB.Model(&model.PageIssue{}).Where("id = ?", id).Group("id DESC").Limit(1).Scan(&jsErrDetail).Error
		err = global.GORMDB.Model(&model.PageIssue{}).Select("id as previous_error_id").Where("id < ? And monitor_id = ? And issues_id = ? ", id, monitorId, jsErrDetail.Issues_id).Group("id DESC").Limit(1).Scan(&jsErrDetail.PreviousErrorID).Error
		err = global.GORMDB.Model(&model.PageIssue{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? And issues_id = ? ", id, monitorId, jsErrDetail.Issues_id).Group("id").Limit(1).Scan(&jsErrDetail.NextErrorID).Error
	}
	return
}
