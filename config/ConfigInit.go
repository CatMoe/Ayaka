package config

import (
	"github.com/CatMoe/Ayaka/logger"
	"github.com/spf13/viper"
)

var (
	PORT    string
	JWTKEY  string
	PREFORK bool
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Panic(err)
	}
	// 加载特定数据
	PORT = viper.GetString("server.port")
	JWTKEY = viper.GetString("server.jwtkey")
	PREFORK = viper.GetBool("server.prefork")
}
