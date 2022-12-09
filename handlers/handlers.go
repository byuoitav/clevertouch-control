package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/clevertouch-control/actions"

	"github.com/gin-gonic/gin"
)

func SetPower(context *gin.Context) {
	power, err := strconv.ParseBool(context.Param("power"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	err = actions.SetPower(context, context.Param("address"), power)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, 1)
}

func GetPower(context *gin.Context) {

	power, err := actions.GetPower(context, context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, power)
}

func SetMute(context *gin.Context) {

	mute, err := strconv.ParseBool(context.Param("mute"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	err = actions.SetMute(context, context.Param("address"), mute)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, 1)
}

func GetMute(context *gin.Context) {

	mute, err := actions.GetMute(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, mute)
}

func SetVolume(context *gin.Context) {

	volume, err := strconv.Atoi(context.Param("volume"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	err = actions.SetVolume(context, context.Param("address"), volume)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, 1)
}

func GetVolume(context *gin.Context) {

	volume, err := actions.GetVolume(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, volume)
}

func SetInput(context *gin.Context) {

	err := actions.SetInput(context, context.Param("address"), context.Param("input"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, 1)
}

func GetInput(context *gin.Context) {

	input, err := actions.GetInput(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, input)
}

func SetBlank(context *gin.Context) {

	blank, err := strconv.ParseBool(context.Param("blank"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	err = actions.SetBlank(context, context.Param("address"), blank)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, 1)
}
