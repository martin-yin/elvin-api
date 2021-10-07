package main

import (
	"dancin-api/core"
	"dancin-api/global"
	"dancin-api/initialize"
)

func main() {
	global.VIPER = core.Viper()       // 初始化Viper
	global.LOGGER = core.Zap()        // 初始化zap日志库
	global.GORMDB = initialize.Gorm() // gorm连接数据库

	if global.GORMDB != nil {
		initialize.MysqlTables(global.GORMDB) // 初始化表
		db, _ := global.GORMDB.DB()
		initialize.Redis()
		defer db.Close()
	}
	go initialize.KafkaWriter()
	go initialize.InitReportData()
	core.RunWindowsServer()
}
