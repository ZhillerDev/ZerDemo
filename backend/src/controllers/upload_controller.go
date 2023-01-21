package controllers

import (
	"ginmod/src/service"
	"github.com/gin-gonic/gin"
)

func UploadSingleImage(ctx *gin.Context) {
	service.UploadSingleImageService(ctx)
}
