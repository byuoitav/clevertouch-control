package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PowerOn(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func Standby(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetPower(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func SwitchInput(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func SetVolume(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func VolumeUnmute(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func VolumeMute(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func setMute(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func BlankDisplay(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func UnblankDisplay(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetVolume(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetInput(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetInputList(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetMute(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GerBlank(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetHardwareInfo(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}

func GetActiveSignal(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
	return
}
