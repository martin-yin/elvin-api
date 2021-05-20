package services

import (
	"danci-api/global"
	"danci-api/model"
	"fmt"
)

func CreatePagePerformance(pagePerformance *model.PagePerformance) {
	if err := global.GVA_DB.Create(&pagePerformance).Error; err != nil {
		fmt.Println("err", err)
	}
	user := model.User{
		PublicFiles: pagePerformance.PublicFiles,
	}
	CreateUser(&user)
}

func CreateUser(user *model.User) {
	// 用户每次刷新页面就记录一次用户。
	if err := global.GVA_DB.Create(&user).Error; err != nil {
		fmt.Println(err, "创建用户出错！")
	}
}

func CreateUserAction(publicFiles model.PublicFiles, reportData string) {
	userAction := model.UserAction{
		UserId:       publicFiles.UserId,
		MonitorId:    publicFiles.MonitorId,
		HappenTime:   publicFiles.HappenTime,
		HappenDay:    publicFiles.HappenDay,
		EventId:      publicFiles.EventId,
		ActionType:   publicFiles.ActionType,
		ActionDetail: reportData,
	}
	global.GVA_DB.Create(&userAction)
}

func CreatePageHttp(pageHttp *model.PageHttp) {
 	global.GVA_DB.Create(&pageHttp)
}

func CreateResourcesError(resourceError *model.PageResourceError) {
	global.GVA_DB.Create(&resourceError)
}

func CreatePageBehavior(pageOperation *model.PageOperation) {
	global.GVA_DB.Create(&pageOperation)
}

func CreatePageJsError(jsError model.PageIssue) {
	global.GVA_DB.Save(&jsError)
}

func CreateJsIssue(stack model.Issue) {
	global.GVA_DB.Save(&stack)
}



func CreatePageView(pageView *model.PageView) {
	global.GVA_DB.Create(&pageView)
}
