package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ConfigurationInit() {
	viper.SetConfigName("config")	// 配置文件名字
	viper.SetConfigType("yaml")		// 配置文件后缀
	viper.AddConfigPath("./")		// 配置文件所在相对路径，路径起始点为项目根目录
	err := viper.ReadInConfig()			// 读入配置
	if err != nil {
		panic(fmt.Errorf("read config err=%s", err))
	} else {
		fmt.Println(viper.GetString("desc"))
	}
}
