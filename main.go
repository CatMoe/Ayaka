package main

import (
	"github.com/CatMoe/Ayaka/config"
	"github.com/CatMoe/Ayaka/handlers"
	"github.com/CatMoe/Ayaka/logger"
	"github.com/gofiber/fiber/v2"
	jwtmiddleware "github.com/gofiber/jwt/v3"
)

func main() {
	// 模块初始化阶段,注意相互依赖顺序
	// 初始化Logger
	logger.LoggerInit()
	// 初始化配置文件
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		Prefork: config.PREFORK,
	})

	api := app.Group("/api/v1")
	// 用户处理模块
	user := api.Group("/user")
	user.Post("/login", handlers.UserLoginHandler)

	// 数据提交处理模块
	//data := api.Group("/data")
	//data.Post("/submit")
	//data.Get("/request/:uid")

	app.Use(jwtmiddleware.New(jwtmiddleware.Config{
		SigningKey: []byte(config.JWTKEY),
	}))

	logger.Logger.Fatal(app.Listen(":" + config.PORT))
}
