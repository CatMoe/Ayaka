package main

import (
	"github.com/CatMoe/Ayaka/config"
	"github.com/CatMoe/Ayaka/handlers"
	"github.com/CatMoe/Ayaka/utils"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	// 模块初始化阶段,注意相互依赖顺序
	// 初始化Utils
	utils.LogUtilInit()

	// 初始化配置文件
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		Prefork: config.PREFORK,
		AppName: "Ayaka",
	})

	// Sentry错误处理中间件
	app.Use(fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: true,
	}))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JWTKEY),
	}))

	// 控制API版本
	api := app.Group("/api/v1")
	// 用户处理模块
	user := api.Group("/user")
	user.Post("/login", handlers.UserLoginHandler)
	user.Post("/register", handlers.UserRegisterHandler)

	logrus.Panic(app.Listen(":" + config.PORT))
}
