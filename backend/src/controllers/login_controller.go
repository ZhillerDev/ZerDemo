package controllers

import (
	"ginmod/src/service"
	"github.com/gin-gonic/gin"
)

func LoginValidate(c *gin.Context) {
	service.LoginValidateService(c)
}

func PageTokenValidate(c *gin.Context) {
	service.PageTokenValidateService(c)
}
