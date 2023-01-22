package main

import (
	"ginmod/config"
	"ginmod/config/logger"
	"ginmod/src/middleware"
	"ginmod/src/preinit"
	r "ginmod/src/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()

	// 1. 初始化服务器配置
	preinit.ServerInit(router)
	preinit.DatabaseInit()

	// 2. 加载配置文件
	config.ConfigurationInit()
	// 3. 加载日志
	logger.LoggerInit()
	logger.Logger.Info("server running")

	// 4. 设置跨域
	router.Use(middleware.CORSSetting())
	// 5. 初始化所有中间件
	router.Use(middleware.Init())

	// 6. 主副路由先后加载
	r.MainRouterInit(router)
	r.TestRouterInit(router)

	// 7. 运行端口
	router.Run(viper.GetString("port"))
}
