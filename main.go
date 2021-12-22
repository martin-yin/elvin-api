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
		// 如果 kafka 连接失败，则使用 redis 消费数据
		if err := initialize.KafkaConn(); err == nil {
			go initialize.ReportDataConsumeByKafka()
		} else {
			go initialize.ReportDataConsumeByRedis()
		}
		defer db.Close()
	}
	core.RunWindowsServer()
}
