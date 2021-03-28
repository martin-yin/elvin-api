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

func WebHttpInfoModel(weLoadPageInfo model.WebHttpInfo, context *gin.Context) {
	err := global.GVA_DB.Create(&weLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func SetWebResourcesError(webResourceErrorInfo model.WebResourceErrorInfo, context *gin.Context) {
	err := global.GVA_DB.Create(&webResourceErrorInfo).Error
	if err != nil {
		fmt.Print(err, "err !!!!!!!! \n")
	}
}


