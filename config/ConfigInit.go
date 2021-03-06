package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	PORT             string
	JWTKEY           string
	PREFORK          bool
	SENTRY_DSN       string
	RECAPTCHA_SECRET string
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Panic(err)
		return
	}
	// 加载特定数据
	PORT = viper.GetString("server.port")
	JWTKEY = viper.GetString("server.jwtkey")
	PREFORK = viper.GetBool("server.prefork")
	SENTRY_DSN = viper.GetString("server.sentry_dsn")
	RECAPTCHA_SECRET = viper.GetString("recaptcha.secret")
}
