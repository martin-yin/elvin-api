package initialize

import (
	"dancin-api/global"
	"dancin-api/middleware"
	"dancin-api/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOGGER.Info("use middleware logger")
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.LOGGER.Info("use middleware cors")
	ApiGroup := Router.Group("")
	router.InitReport(ApiGroup)   // 上报数据
	router.InitCommunal(ApiGroup) // 公用接口
	router.InitAdmin(ApiGroup)    // 管理相关接口
	global.LOGGER.Info("router register success")
	return Router
}
