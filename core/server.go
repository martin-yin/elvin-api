package core

import (
	"dancin-api/global"
	"dancin-api/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONFIG.System.UseMultipoint {
		initialize.Redis()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Info("server run success on ", zap.String("address", address))
	global.LOGGER.Error(s.ListenAndServe().Error())
}
