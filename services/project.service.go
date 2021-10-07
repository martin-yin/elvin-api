package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/response"
	"fmt"
	"strconv"
	"strings"
)

func GetProjectList(id uint) (projectList []model.Project, err error) {
	var teamList []model.Team
	err = global.GORMDB.Preload("Admins", "id = ? ", 1).Preload("Projects").Model(&model.Team{}).Find(&teamList).Error
	for _, team := range teamList {
		for _, project := range team.Projects {
			projectList = append(projectList, project)
		}
	}
	return
}

func CreateProject(project model.Project) (err error) {
	err = global.GORMDB.Create(&project).Error
	return err
}

func FindProject(projectName string) (isExist bool) {
	var project model.Project
	result := global.GORMDB.Model(&model.Project{}).Where("project_name = ? ", projectName).First(&project)
	if result.RowsAffected != 0 {
		return true
	}
	return false
}

func GetProject(monitorId string) (project model.Project, err error) {
	err = global.GORMDB.Model(&model.Project{}).Where("monitor_id = ? ", monitorId).First(&project).Error
	return
}

func GetProjectStatistics(startTime string, endTime string, monitorId string) (projectStatistics response.ProjectStatistics, err error) {
	err = global.GORMDB.Model(&model.PageView{}).Select("COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(SqlWhereBuild("page_views"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
	return
}

func DelProject(id string) (err error) {
	err = global.GORMDB.Delete(&model.Project{}, id).Error
	return err
}

func GetProjectHealthy(startTime string, endTime string, monitorIds string) (projectStatisticsList []response.ProjectStatistics, err error) {
	monitorIdealists := strings.Split(monitorIds, `,`)
	for _, monitorId := range monitorIdealists {
		var projectStatistics response.ProjectStatistics
		err = global.GORMDB.Model(&model.PageView{}).Select("COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(SqlWhereBuild("page_views"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		err = global.GORMDB.Model(&model.Issue{}).Select("COUNT( DISTINCT id ) as js_error").Where(SqlWhereBuild("issues"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		err = global.GORMDB.Model(&model.PageResourceError{}).Select("COUNT( DISTINCT id ) as resources_error").Where(SqlWhereBuild("page_resource_errors"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		err = global.GORMDB.Model(&model.PageHttp{}).Select("COUNT( DISTINCT id ) as http_error").Where(SqlWhereBuild("page_https"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		projectStatistics.JsError = DecimalNotZero(projectStatistics.JsError, projectStatistics.Pv)
		projectStatistics.ResourcesError = DecimalNotZero(projectStatistics.ResourcesError, projectStatistics.Pv)
		projectStatistics.HttpError = DecimalNotZero(projectStatistics.HttpError, projectStatistics.Pv)
		projectStatisticsList = append(projectStatisticsList, projectStatistics)
	}
	return
}

func DecimalNotZero(value float64, value2 float64) float64 {
	if value != 0 && value2 != 0 {
		result := value / value2
		result, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
		return result
	}
	return 0
}

func SqlWhereBuild(model string) string {
	return "from_unixtime(" + model + ".happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? "
}
