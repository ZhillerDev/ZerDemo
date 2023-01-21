package service

import (
	"ginmod/src/constants"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path"
	"strings"
)

func UploadSingleImageService(ctx *gin.Context) {
	img, err := ctx.FormFile("file")
	if err != nil {
		BasicResponseService(
			ctx,
			http.StatusBadRequest,
			constants.GetMessage(constants.ERROR_UPLOAD_SAVE_IMAGE))
		return
	}

	suffix := strings.ToLower(path.Ext(img.Filename))
	if allowSuffix := ".jpg.png.jpeg.gif"; !strings.Contains(allowSuffix, suffix) {
		BasicResponseService(
			ctx,
			http.StatusBadRequest,
			constants.GetMessage(constants.ERROR_UPLOAD_TYPE_IMAGE))
		return
	}

	filePath := viper.GetString("image-save-path")
	_, err2 := os.Stat(filePath + "/single")
	if err2 != nil {
		os.Mkdir(filePath+"/single", os.ModePerm)
	}

	fileName := filePath + "/single/demo.jpg"
	ctx.SaveUploadedFile(img, fileName)
	BasicResponseService(
		ctx,
		http.StatusOK,
		viper.GetString("static-mainurl")+"/uploads/images/single/demo.jpg")
}
