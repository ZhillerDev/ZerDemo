package router

import (
	"ginmod/src/controllers"
	"github.com/gin-gonic/gin"
)

func MainRouterInit(engine *gin.Engine) {
	// 文件及图片上传路由
	uploadsRouter := engine.Group("/uploads")
	{
		uploadsRouter.POST("/image", controllers.UploadSingleImage)
	}

	// 登录与注册路由
	loginRouter := engine.Group("/login")
	{
		loginRouter.POST("/validate", controllers.LoginValidate)
		loginRouter.POST("/check", controllers.PageTokenValidate)
	}

	// 数据库查询路由
	selectDatabaseRouter := engine.Group("/sdb")
	{
		selectDatabaseRouter.GET("/allusers", controllers.SelectAllUsers)
	}
}
