package services

import (
	"danci-api/global"
	"danci-api/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetWebPerformance(pagePerformance model.WebPerformance, context *gin.Context) {
	err := global.GVA_DB.Create(&pagePerformance).Error
	fmt.Print(pagePerformance.ID)
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func SetWebRequest(request model.WebRequest, context *gin.Context) {
	err := global.GVA_DB.Create(&request).Error
	if err != nil {
		fmt.Print(err, "err !!!!!!!! \n")
	}
}

func SetWebResourcesError(resourceError model.WebResourcesError, context *gin.Context) {
	err := global.GVA_DB.Create(&resourceError).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func GetWebResourcesErrorCount() (userCount int64, errCount int64, pageViewCount int64) {
	errorBb := global.GVA_DB.Model(&model.WebResourcesError{})
	db := global.GVA_DB.Model(&model.WebPerformance{})
	// 总用户数量  UV
	errorBb.Group("cookie").Count(&userCount)
	// 总资源错误次数
	errorBb.Count(&errCount)
	// 总PV 数量
	db.Count(&pageViewCount)
	return userCount, errCount, pageViewCount
}
