package router

import (
	"ginmod/src/controllers"
	"github.com/gin-gonic/gin"
)

func TestRouterInit(engine *gin.Engine) {
	testRouter := engine.Group("/test")
	{
		testRouter.GET("/getstr", controllers.GetServerPort)
	}
}
