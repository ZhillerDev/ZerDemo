package middleware

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gin-gonic/gin"
	"time"
)

func Init() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("开始拦截 " + convertor.ToString(time.Now().Unix()))

		// JWT拦截初始化

	}
}
