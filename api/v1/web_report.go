package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 往数据中添加web性能
func SetWebLoadPageInfo(context *gin.Context) {
	var webLoadPageInfo request.WebLoadPageInfo
	err := context.BindJSON(&webLoadPageInfo)
	if err != nil {
		fmt.Print(err)
	}
	pagePerformanceModel := &model.WebLoadpageInfo{
		//PageUrl string `json:"page_url"`
		//UserId      int    `json:"user_id"`
		//UploadType  int    `json:"upload_type"`
		//HappenTime  int    `json:"happen_time"`
		//HappenDate  int    `json:"happen_date"`
		//
		//PageKey     string `json:"page_key"`
		PageUrl: webLoadPageInfo.PageUrl,
		UserId: "",
		UploadType: "LOAD_PAGE",
		HappenTime: "",
		HappenDate: "",
		PageKey: webLoadPageInfo.PageKey,

		LoadPage: webLoadPageInfo.LoadPage,
		DomReady: webLoadPageInfo.DomReady,
		Redirect: webLoadPageInfo.Redirect,

		LookupDomain: webLoadPageInfo.LookupDomain,
		Ttfb: webLoadPageInfo.Ttfb,
		Request: webLoadPageInfo.Request,
		LoadEvent: webLoadPageInfo.LoadEvent,
		Appcache: webLoadPageInfo.Appcache,
		UnloadEvent: webLoadPageInfo.UnloadEvent,

		Connect: webLoadPageInfo.Connect,
		LoadType: webLoadPageInfo.LoadType,
		BrowserInfo: webLoadPageInfo.BrowserInfo,
	}
	services.SetWebLoadPageInfo(*pagePerformanceModel, context)
}

// 存储页面得请求数据
func SetWebRequest(context *gin.Context) {
	//var pageRequest request.PageRequest
	//err := context.BindJSON(&pageRequest)
	//if err != nil {
	//
	//	//global.GVA_LOG.Error(err)
	//}
	//cookie, err := context.Cookie("gin_cookie")
	//pageRequestModel := &model.WebRequest{
	//	GenerateTime: pageRequest.Data.GenerateTime,
	//	Method:       pageRequest.Data.Method,
	//	HttpType:     pageRequest.Data.HttpType,
	//	ElapsedTime:  pageRequest.Data.ElapsedTime,
	//	Code:         pageRequest.Data.Code,
	//	ApiUrl:       pageRequest.Data.ApiUrl,
	//	PageUrl:      pageRequest.Data.PageUrl,
	//	Message:      pageRequest.Data.Message,
	//	IsError:      pageRequest.Data.IsError,
	//
	//	UserAgent: context.Request.UserAgent(),
	//	IPAddress: context.ClientIP(),
	//}
	//services.SetWebRequest(*pageRequestModel, context)
}

func SetWebResourcesError(context *gin.Context) {
	//var pageResourcesError request.PageResourcesError
	//err := context.BindJSON(&pageResourcesError)
	//if err != nil {
	//	fmt.Print(err, "err!")
	//}
	//context.Cookie()
	//pageResourcesErrorModel := &model.WebResourcesError{
	//	PageUrl:      pageResourcesError.Data.PageUrl,
	//	ResourceUrl:  pageResourcesError.Data.ResourceUrl,
	//	GenerateTime: pageResourcesError.Data.GenerateTime,
	//	DomPath:      pageResourcesError.Data.DomPath,
	//
	//	UserAgent: context.Request.UserAgent(),
	//	IPAddress: context.ClientIP(),
	//}
	//services.SetWebResourcesError(*pageResourcesErrorModel, context)
}

// 将uuid 写入cookie， 一个cookie 就是一个用户。每次切换一个页面等页面加载完成的时候，就是上报一次pv。
//  cookie 数量 = 用户数量
// 记录每次报错来自哪个cookie， 然后连表分析 报错影响的 用户数量 和 PV
// 计算一分钟之内的影响 =  一分钟之内有多少个报错   一分钟之内有多少个PV 一分钟之内有多少个用户
// 获取apikey 获取对应项目得资源错误，计算资源异常PV 次数 和 影响用户数量
func GetWebResourcesError(context *gin.Context) {

	//var userCount, errCount, pageViewCount = services.GetWebResourcesErrorCount();

	context.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   nil,
	})
}

// 获取apikey 获取对应项目得JS错误，计算资源异常PV 次数 和 影响用户数量
func GetWebJsError(context *gin.Context) {
}
