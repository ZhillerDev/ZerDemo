package router

import (
	"ginmod/src/controllers"
	"github.com/gin-gonic/gin"
)

func MainRouterInit(engine *gin.Engine) {
	uploadsRouter := engine.Group("/uploads")
	{
		uploadsRouter.POST("/image", controllers.UploadSingleImage)
	}
}
