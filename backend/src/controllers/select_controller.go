package controllers

import (
	"ginmod/src/service"
	"github.com/gin-gonic/gin"
)

func SelectAllUsers(c *gin.Context)  {
	service.SelectAllUsersService(c)
}
