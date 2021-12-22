package utils

import (
	"dancin-api/global"
	"dancin-api/model"
	"encoding/json"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"go.uber.org/zap"
	"os"
)
var Ip2region *ip2region.Ip2Region
var IPChan = make(chan string)

func init()  {
	Ip2region = NewIp2Region()
}

func NewIp2Region() *ip2region.Ip2Region {
	basePath, err := os.Getwd()
	region, _ := ip2region.New(basePath + "/ip2region.db")
	if err != nil {
		panic(err)
	}
	defer region.Close()
	return region
}

func ConsumeIP() {
	for ip := range IPChan {
		transformIPToAddress(ip)
	}
}

func transformIPToAddress(ip string) {
	addingStr := global.REDIS.HGet("ipAddress", ip)
	if len(addingStr.Val()) != 0 {
		return
	} else {
		ipInfo, err := Ip2region.BtreeSearch(ip)
		global.LOGGER.Error("ip2region ip解析出错:", zap.Any("err", err))
		ipInfoStr, err := json.Marshal(ipInfo)
		global.LOGGER.Error("ip2region ip解析出错:", zap.Any("err", err))
		global.REDIS.HSet("ipAddress", ip, ipInfoStr)
		// 在数据中写一份数据。
		ipAddressData := &model.IPAddress{
			IP:       ip,
			ISP:      ipInfo.ISP,
			CityId:   ipInfo.CityId,
			Country:  ipInfo.Country,
			Region:   ipInfo.Region,
			Province: ipInfo.Province,
			City:     ipInfo.City,
		}
		if err != nil {
			global.LOGGER.Error("ip2region ip解析出错:", zap.Any("err", err))
			return
		}
		err = global.GORMDB.Model(&model.IPAddress{}).Create(ipAddressData).Error
		global.LOGGER.Error("ip地址写入数据库出错:", zap.Any("err", err))
	}
}
