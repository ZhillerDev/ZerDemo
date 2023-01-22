package service

import (
	"fmt"
	"ginmod/src/constants"
	"ginmod/src/middleware"
	"ginmod/src/model"
	"ginmod/src/preinit"
	"github.com/gin-gonic/gin"
)

func LoginValidateService(c *gin.Context) {

	if originToken := c.Request.Header.Get("token"); originToken != "" {
		if middleware.ValidateToken(originToken) {
			constants.RespJSONMessageSuccess(c, originToken)
			return
		}
	}

	var user *model.User

	fmt.Println(c.Request)

	username := c.PostForm("username")
	password := c.PostForm("password")

	preinit.DB.Where("username = ? AND password = ?", username, password).Find(&user)

	fmt.Println(user)

	if user.Username == "" {
		constants.RespJSONMessageFailed(c, "数据库校验错误，登录失败！")
		return
	}

	token, _ := middleware.GenerateToken(username)
	constants.RespJSONMessageSuccess(c, token)
}

func PageTokenValidateService(c *gin.Context) {
	token := c.PostForm("token")
	res := middleware.ValidateToken(token)
	if !res{
		constants.RespJSONMessageFailed(c,"token已经失效！")
		return
	}
	constants.RespJSONMessageSuccess(c,"token可以使用！")
}
