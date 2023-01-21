package service

import "github.com/gin-gonic/gin"

func BasicResponseService(ctx *gin.Context, codeId int, msg string) {
	ctx.JSON(codeId, gin.H{
		"code": codeId,
		"msg":  msg,
	})
}


