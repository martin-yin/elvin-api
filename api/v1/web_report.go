package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetWebPerformance(context *gin.Context) {
	var pagePerformance request.PagePerformance
	err := context.BindJSON(&pagePerformance)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": 200,
			"error":  err,
			"data":   pagePerformance,
		})
	}

	cookie, err := context.Cookie("gin_cookie")
	pagePerformanceModel := &model.WebPerformance{
		Appcache:        pagePerformance.Data.Appcache,
		Contentdownload: pagePerformance.Data.Contentdownload,
		Dns:             pagePerformance.Data.Dns,
		Domparsing:      pagePerformance.Data.Domparsing,
		PageUrl:         pagePerformance.Data.PageUrl,
		Redirect:        pagePerformance.Data.Redirect,
		Res:             pagePerformance.Data.Res,
		Tcp:             pagePerformance.Data.Tcp,
		Ttfb:            pagePerformance.Data.Ttfb,
		UserAgent:       context.Request.UserAgent(),
		IPAddress:       context.ClientIP(),
		Cooike:          cookie,
	}
	services.SetWebPerformance(*pagePerformanceModel, context)
}

// 存储页面得请求数据
func SetWebRequest(context *gin.Context) {
	var pageRequest request.PageRequest
	err := context.BindJSON(&pageRequest)
	if err != nil {
		fmt.Print(err, "err!")
		//global.GVA_LOG.Error(err)
	}
	//cookie, err := context.Cookie("gin_cookie")
	pageRequestModel := &model.WebRequest{
		GenerateTime: pageRequest.Data.GenerateTime,
		Method:       pageRequest.Data.Method,
		HttpType:     pageRequest.Data.HttpType,
		ElapsedTime:  pageRequest.Data.ElapsedTime,
		Code:         pageRequest.Data.Code,
		ApiUrl:       pageRequest.Data.ApiUrl,
		PageUrl:      pageRequest.Data.PageUrl,
		Message:      pageRequest.Data.Message,
		IsError:      pageRequest.Data.IsError,

		UserAgent: context.Request.UserAgent(),
		IPAddress: context.ClientIP(),
	}
	services.SetWebRequest(*pageRequestModel, context)
}

func SetWebResourcesError(context *gin.Context) {
	var pageResourcesError request.PageResourcesError
	err := context.BindJSON(&pageResourcesError)
	if err != nil {
		fmt.Print(err, "err!")
	}
	//context.Cookie()
	pageResourcesErrorModel := &model.WebResourcesError{
		PageUrl:      pageResourcesError.Data.PageUrl,
		ResourceUrl:  pageResourcesError.Data.ResourceUrl,
		GenerateTime: pageResourcesError.Data.GenerateTime,
		DomPath:      pageResourcesError.Data.DomPath,

		UserAgent: context.Request.UserAgent(),
		IPAddress: context.ClientIP(),
	}
	services.SetWebResourcesError(*pageResourcesErrorModel, context)
}

// 将uuid 写入cookie， 一个cookie 就是一个用户。每次切换一个页面等页面加载完成的时候，就是上报一次pv。
//  cookie 数量 = 用户数量
// 记录每次报错来自哪个cookie， 然后连表分析 报错影响的 用户数量 和 PV
// 计算一分钟之内的影响 =  一分钟之内有多少个报错   一分钟之内有多少个PV 一分钟之内有多少个用户
// 获取apikey 获取对应项目得资源错误，计算资源异常PV 次数 和 影响用户数量
func GetWebResourcesError(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   services.GetWebResourcesErrorCount(),
	})
}

// 获取apikey 获取对应项目得JS错误，计算资源异常PV 次数 和 影响用户数量
func GetWebJsError(context *gin.Context) {
}
