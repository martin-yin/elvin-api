package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/utils"
)

func FindIssue(stack string) (issuesModel model.Issue, err error) {
	err = global.GORMDB.Where("stack = ? ", stack).Find(&issuesModel).Error
	return
}

func GetIssues(params request.RequestParams) (issues []response.Issues, err error) {
	var issueList []response.Issues
	sql, sqlParams := utils.BuildWhereSql("page_js_errors", "", params)
	err = global.GORMDB.Model(&model.Issue{}).Select("issues.id, "+
		"issues.error_name, "+
		"issues.message, "+
		"page_js_errors.happen_time, page_js_errors.monitor_id, "+
		"COUNT(DISTINCT page_js_errors.user_id) as error_user, "+
		"COUNT(page_js_errors.id) as total").Joins(""+
		"INNER JOIN page_js_errors on page_js_errors.issues_id = issues.id"+
		"").Group("issues.id").Where(sql, sqlParams...).Debug().Find(&issueList).Error

	for _, issue := range issueList {
		err = global.GORMDB.Model(&model.PageJsError{}).Select("happen_time as first_time").Where("issues_id = ? ", &issue.ID).Group("id ASC").Limit(1).Scan(&issue).Error
		err = global.GORMDB.Model(&model.PageJsError{}).Select("happen_time as last_time").Where("issues_id = ? ", &issue.ID).Limit(1).Scan(&issue).Error
		issues = append(issues, issue)
	}
	return
}

func GetJsErrorDetail(id, errId int, monitorId string) (jsErrorDetail response.PageJsErrorDetail, err error) {
	if errId != 0 && id == 0 {
		err = global.GORMDB.Model(&model.PageJsError{}).Where("issues_id = ? And monitor_id = ?", errId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrorDetail).Error
		err = global.GORMDB.Model(&model.PageJsError{}).Select("id as previous_error_id ").Where("id < ? And monitor_id = ? And issues_id = ? ", jsErrorDetail.ID, monitorId, errId).Group("id DESC").Limit(1).Scan(&jsErrorDetail.PreviousErrorID).Error
		err = global.GORMDB.Model(&model.PageJsError{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? And issues_id = ? ", jsErrorDetail.ID, monitorId, errId).Group("id").Limit(1).Scan(&jsErrorDetail.NextErrorID).Error
	} else if id != 0 {
		err = global.GORMDB.Model(&model.PageJsError{}).Where("id = ?", id).Group("id DESC").Limit(1).Scan(&jsErrorDetail).Error
		err = global.GORMDB.Model(&model.PageJsError{}).Select("id as previous_error_id").Where("id < ? And monitor_id = ? And issues_id = ? ", id, monitorId, jsErrorDetail.Issues_id).Group("id DESC").Limit(1).Scan(&jsErrorDetail.PreviousErrorID).Error
		err = global.GORMDB.Model(&model.PageJsError{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? And issues_id = ? ", id, monitorId, jsErrorDetail.Issues_id).Group("id").Limit(1).Scan(&jsErrorDetail.NextErrorID).Error
	}
	return
}
