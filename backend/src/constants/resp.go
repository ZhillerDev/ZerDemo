package constants

import "github.com/gin-gonic/gin"

func RespJSONMessageSuccess(c *gin.Context, msg string) {
	c.JSON(SUCCESS, gin.H{
		"code": SUCCESS,
		"msg":  msg,
	})
}

func RespJSONMessageFailed(c *gin.Context, msg string) {
	c.JSON(FAILED, gin.H{
		"code": FAILED,
		"msg":  msg,
	})
}

func RespJSONMessageError(c *gin.Context, msg string) {
	c.JSON(ERROR, gin.H{
		"code": ERROR,
		"msg":  msg,
	})
}

func RespJSONData(c *gin.Context, datas any) {
	c.JSON(SUCCESS, gin.H{
		"status": SUCCESS,
		"data":   datas,
	})
}
