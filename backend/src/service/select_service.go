package service

import (
	"ginmod/src/constants"
	"ginmod/src/model"
	"ginmod/src/preinit"
	"github.com/gin-gonic/gin"
)

func SelectAllUsersService(c *gin.Context) {
	var users []model.User
	preinit.DB.Find(&users)
	constants.RespJSONData(c, users)
}
