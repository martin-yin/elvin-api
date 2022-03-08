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
	sql, sqlParams := utils.BuildWhereSql("page_js_errs", "", params)
	err = global.GORMDB.Model(&model.Issue{}).Select("issues.id, "+
		"issues.error_name, "+
		"issues.message, "+
		"page_js_errs.happen_time, page_js_errs.monitor_id, "+
		"COUNT(DISTINCT page_js_errs.user_id) as error_user, "+
		"COUNT(page_js_errs.id) as total").Joins(""+
		"INNER JOIN page_js_errs on page_js_errs.issues_id = issues.id"+
		"").Group("issues.id").Where(sql, sqlParams...).Debug().Find(&issueList).Error

	for _, issue := range issueList {
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("page_url, created_at as first_time").Where("issues_id = ? ", &issue.ID).Group("id ASC").Limit(1).Scan(&issue).Error
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("created_at as last_time").Where("issues_id = ? ", &issue.ID).Group("id DESC").Limit(1).Scan(&issue).Error
		// Todo  感觉不需要了，可以考虑删掉。
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("COUNT(id) as today").Where("from_unixtime(page_js_errs.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", params.StartTime, params.EndTime).Scan(&issue).Error
		issues = append(issues, issue)
	}
	return
}

func GetIssuesDetail(id, errId int, monitorId string) (jsErrDetail response.PageJsErrDetail, err error) {
	if errId != 0 && id == 0 {
		err = global.GORMDB.Model(&model.PageJsErr{}).Where("issues_id = ? And monitor_id = ?", errId, monitorId).Group("id DESC").Limit(1).Scan(&jsErrDetail).Error
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("id as previous_error_id ").Where("id < ? And monitor_id = ? And issues_id = ? ", jsErrDetail.ID, monitorId, errId).Group("id DESC").Limit(1).Scan(&jsErrDetail.PreviousErrorID).Error
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? And issues_id = ? ", jsErrDetail.ID, monitorId, errId).Group("id").Limit(1).Scan(&jsErrDetail.NextErrorID).Error
	} else if id != 0 {
		err = global.GORMDB.Model(&model.PageJsErr{}).Where("id = ?", id).Group("id DESC").Limit(1).Scan(&jsErrDetail).Error
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("id as previous_error_id").Where("id < ? And monitor_id = ? And issues_id = ? ", id, monitorId, jsErrDetail.Issues_id).Group("id DESC").Limit(1).Scan(&jsErrDetail.PreviousErrorID).Error
		err = global.GORMDB.Model(&model.PageJsErr{}).Select("id as next_error_id").Where("id > ? And monitor_id = ? And issues_id = ? ", id, monitorId, jsErrDetail.Issues_id).Group("id").Limit(1).Scan(&jsErrDetail.NextErrorID).Error
	}
	return
}
