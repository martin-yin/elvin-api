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
		UserId:      pagePerformance.PublicFiles.UserId,
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
	userAciton := model.UserAction{
		UserId:       publicFiles.UserId,
		MonitorId:    publicFiles.MonitorId,
		HappenTime:   publicFiles.HappenTime,
		HappenDay:    publicFiles.HappenDay,
		EventId:      publicFiles.EventId,
		ActionType:   publicFiles.ActionType,
		ActionDetail: reportData,
	}
	if err := global.GVA_DB.Create(&userAciton).Error; err != nil {
		fmt.Println(err, "err \n")
	}
}

func CreatePageHttp(pageHttp *model.PageHttp) {
	if err := global.GVA_DB.Create(&pageHttp).Error; err != nil {
		fmt.Println(err, "err \n")
	}
}

func CreateResourcesError(resourceError *model.PageResourceError) {
	if err := global.GVA_DB.Create(&resourceError).Error; err != nil {
		fmt.Println(err, "err \n")
	}
}

func CreatePageBehavior(pageOperation *model.PageOperation) {
	if err := global.GVA_DB.Create(&pageOperation).Error; err != nil {
		fmt.Println(err, "err \n")
	}
}

func CreatePageJsError(pageJsError *model.PageJsError) {
	if err := global.GVA_DB.Create(&pageJsError).Error; err != nil {
		fmt.Println(err, "err \n")
	}
}

func CreatePageView(pageView *model.PageView) {
	if err := global.GVA_DB.Create(&pageView).Error; err != nil {
		fmt.Println(err, "err \n")
	}
}
