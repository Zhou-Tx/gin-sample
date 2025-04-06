package main

import (
	"fmt"
	"gin-sample/config"
	"gin-sample/controller"
	"gin-sample/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	server := config.LoadConfig().Server

	// 初始化 Gin 引擎
	engine := gin.Default()
	// 注册全局中间件
	engine.Use(middleware.AuthMiddleware())
	// 注册路由
	controller.RegisterRoutes(engine)

	// 启动服务
	_ = engine.Run(fmt.Sprintf("%s:%d", server.Host, server.Port))
}
