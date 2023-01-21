package main

import (
	"ginmod/config"
	"ginmod/config/logger"
	"ginmod/src/middleware"
	r "ginmod/src/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func PreInit(engine *gin.Engine) {
	engine.Static("/static", "./runtime")
}

func main() {
	router := gin.Default()

	// 预初始化：路由自身的一些简单性质
	PreInit(router)

	// 加载配置文件
	config.ConfigurationInit()

	// 加载日志
	logger.LoggerInit()
	logger.Logger.Info("server running")

	// 设置跨域
	router.Use(middleware.CORSSetting())

	// 主副路由先后加载
	r.MainRouterInit(router)
	r.TestRouterInit(router)

	// 运行端口
	router.Run(viper.GetString("port"))
}
