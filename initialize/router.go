package initialize

import (
	"danci-api/global"
	"danci-api/middleware"
	"danci-api/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	//global.GVA_LOG.Info("use middleware cors")
	ApiGroup := Router.Group("")
	router.InitReport(ApiGroup)  // 上报数据
	router.InitCommunal(ApiGroup)  // 公用接口
	router.InitAdmin(ApiGroup) // 管理相关接口
	global.GVA_LOG.Info("router register success")
	return Router
}
