package controllers

import (
	"ginmod/src/constants"
	"ginmod/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServerPort(ctx *gin.Context) {
	service.BasicResponseService(
		ctx,
		http.StatusOK,
		constants.GetMessage(constants.SUCCESS))
}
