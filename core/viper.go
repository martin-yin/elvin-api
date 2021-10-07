package core

import (
	"dancin-api/global"
	"dancin-api/utils"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Viper() *viper.Viper {
	var config string
	if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
		config = utils.ConfigFile
		fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
