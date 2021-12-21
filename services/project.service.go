package services

import (
	"dancin-api/global"
	"dancin-api/model"
	"dancin-api/model/response"
	"fmt"
	"strconv"
)

func GetProjectList(id uint) (projectList []model.Project, err error) {
	var teamList []model.Team
	err = global.GORMDB.Preload("Admins", "id = ? ", id).Preload("Projects").Model(&model.Team{}).Find(&teamList).Error
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

func DelProject(id string) (err error) {
	// todo 判断是否能删除
	err = global.GORMDB.Delete(&model.Project{}, id).Error
	return err
}

func GetHealthStatus(id uint, startTime string, endTime string) (homeStatistics []response.HomeStatistic, err error) {
	projectList, err := GetProjectList(id)
	if err != nil {
		return nil, err
	}
	// 获取当前用户下的项目
	for _, project := range projectList {
		var homeStatistic response.HomeStatistic
		homeStatistic.ProjectName = project.ProjectName
		homeStatistic.MonitorId = project.MonitorId
		err = global.GORMDB.Model(&model.PageView{}).Select("COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(SqlWhereBuild("page_views"), startTime, endTime, project.MonitorId).Scan(&homeStatistic).Error
		err = global.GORMDB.Model(&model.Issue{}).Select("COUNT( DISTINCT id ) as js_error").Where(SqlWhereBuild("issues"), startTime, endTime, project.MonitorId).Scan(&homeStatistic).Error
		err = global.GORMDB.Model(&model.PageResourceError{}).Select("COUNT( DISTINCT id ) as resources_error").Where(SqlWhereBuild("page_resource_errors"), startTime, endTime, project.MonitorId).Scan(&homeStatistic).Error
		err = global.GORMDB.Model(&model.PageHttp{}).Select("COUNT( DISTINCT id ) as http_error").Where(SqlWhereBuild("page_https"), startTime, endTime, project.MonitorId).Scan(&homeStatistic).Error
		homeStatistic.JsError = DecimalNotZero(homeStatistic.JsError, homeStatistic.Pv)
		homeStatistic.ResourcesError = DecimalNotZero(homeStatistic.ResourcesError, homeStatistic.Pv)
		homeStatistic.HttpError = DecimalNotZero(homeStatistic.HttpError, homeStatistic.Pv)
		homeStatistics = append(homeStatistics, homeStatistic)
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
	return "from_unixtime(" + model + ".happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and " + model + ".monitor_id = ? "
}


func SqlBuildWhere(model, sql string, value ...interface{}) string {
	fmt.Print(value, "===================")
	return "from_unixtime(" + model + ".happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and " + model + ".monitor_id = ? " + sql
}

func TestMain() {
	type test struct {
		StartTime string
		EndTime string
	}
	cc := &test{
		StartTime: "",
		EndTime: "",
	}
	SqlBuildWhere("", "",cc)
}
