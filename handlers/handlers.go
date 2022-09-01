package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/byuoitav/clevertouch-control/actions"

	"github.com/gin-gonic/gin"
)

func PowerOn(context *gin.Context) {

	err := actions.SetPower(context, context.Param("address"), true)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func Standby(context *gin.Context) {

	err := actions.SetPower(context, context.Param("address"), false)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, 1)
}

func GetPower(context *gin.Context) {

	response, err := actions.GetPower(context, context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, response)
}

func SwitchInput(context *gin.Context) {
	address := context.Param("address")
	port := context.Param("port")

	splitPort := strings.Split(port, "!")

	params := make(map[string]interface{})
	if len(splitPort) < 2 {
		context.JSON(http.StatusBadRequest, fmt.Sprintf("ports cofigured incorrectly (should follow format \"hdmi!2\"): %s", port))
	}
	params["uri"] = fmt.Sprintf("extInput:%s?port=%s", splitPort[0], splitPort[1])

	err := actions.BuildAndSendPayload(address, "avContent", "setPlayContent", params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func SetVolume(context *gin.Context) {
	address := context.Param("address")
	value := context.Param("value")

	volume, err := strconv.Atoi(value)
	if err != nil {
		context.JSON(http.StatusBadRequest, fmt.Sprintf("volume value must be an integer: %s", value))
	} else if volume < 0 || volume > 100 {
		context.JSON(http.StatusBadRequest, fmt.Sprintf("volume value must be between 0 and 100: %s", value))
	}

	params := make(map[string]interface{})
	params["target"] = "speaker"
	params["volume"] = volume

	err = actions.BuildAndSendPayload(address, "audio", "setAudioVolume", params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	params = make(map[string]interface{})
	params["target"] = "headphone"
	params["volume"] = volume

	err = actions.BuildAndSendPayload(address, "audio", "setAudioVolume", params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func VolumeUnmute(context *gin.Context) {

	address := context.Param("address")

	err := setMute(context, address, false, 4)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func setMute(context *gin.Context, address string, status bool, retryCount int) error {
	params := make(map[string]interface{})
	params["status"] = status

	initCount := retryCount

	for retryCount >= 0 {
		err := actions.BuildAndSendPayload(address, "audio", "setAudioMute", params)
		if err != nil {
			return err
		}

		postStauts, err := actions.GetMute(address)
		if err != nil {
			return err
		}

		if true == status {
			return nil
		}

		retryCount--

		time.Sleep(10 * time.Millisecond)
	}

	return fmt.Errorf("Attempted to set mute status %v times, could not set mute status", initCount+1)
}

func VolumeMute(context *gin.Context) {

	err := setMute(context, context.Param("address"), true, 4)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func BlankDisplay(context *gin.Context) {
	params := make(map[string]interface{})
	params["mode"] = "pictureOff"

	err := actions.BuildAndSendPayload(context.Param("address"), "system", "setPowerSavingMode", params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func UnblankDisplay(context *gin.Context) {
	params := make(map[string]interface{})
	params["mode"] = "off"

	err := actions.BuildAndSendPayload(context.Param("address"), "system", "setPowerSavingMode", params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, 1)
}

func GetVolume(context *gin.Context) {
	response, err := actions.GetVolume(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, response)
}

func GetInput(context *gin.Context) {
	response, err := actions.GetInput(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, response)
}

func GetInputList(context *gin.Context) {

	context.JSON(http.StatusOK, 1)
}

func GetMute(context *gin.Context) {
	response, err := actions.GetMute(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, response)
}

func GetBlank(context *gin.Context) {
	response, err := actions.GetBlanked(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, response)
}

func GetHardwareInfo(context *gin.Context) {
	response, err := actions.GetHardwareInfo(context.Param("address"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, response)
}

func GetActiveSignal(context *gin.Context) {
	response, err := actions.GetActiveSignal(context.Param("port"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, response)
}
