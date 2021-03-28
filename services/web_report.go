package services

import (
	"danci-api/global"
	"danci-api/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetWebLoadPageInfo(weLoadPageInfo model.WebLoadpageInfo, context *gin.Context) {
	err := global.GVA_DB.Create(&weLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

//func SetWebPerformance(pagePerformance model.WebPerformance, context *gin.Context) {
//
//}
//
//func SetWebRequest(request model.WebRequest, context *gin.Context) {
//	err := global.GVA_DB.Create(&request).Error
//	if err != nil {
//		fmt.Print(err, "err !!!!!!!! \n")
//	}
//}
//
//func SetWebResourcesError(resourceError model.WebResourcesError, context *gin.Context) {
//	err := global.GVA_DB.Create(&resourceError).Error
//	if err != nil {
//		fmt.Print(err, "err!!!!!!!!!!!!! \n")
//	}
//}

