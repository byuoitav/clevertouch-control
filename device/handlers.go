package device

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/clevertouch-control/device/actions"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func (d *DeviceManager) setPower(context *gin.Context) {
	d.Log.Debug("setting power", zap.String("power", context.Param("power")), zap.String("address", context.Param("address")))
	power, err := strconv.ParseBool(context.Param("power"))
	if err != nil {
		d.Log.Warn("could not set power. 'power' parameter not a valid boolean value", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = actions.SetPower(context, context.Param("address"), power)
	if err != nil {
		d.Log.Warn("failed to set power", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("successfully set power", zap.String("power", context.Param("power")), zap.String("address", context.Param("address")))
	context.JSON(http.StatusOK, 1)
}

func (d *DeviceManager) getPower(context *gin.Context) {
	d.Log.Debug("getting power status", zap.String("address", context.Param("address")))

	power, err := actions.GetPower(context, context.Param("address"))
	if err != nil {
		d.Log.Warn("could not get power", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("received power status", zap.String("power", power.Power))
	context.JSON(http.StatusOK, power)
}

func (d *DeviceManager) getBooted(context *gin.Context) {
	d.Log.Debug("getting booted status", zap.String("address", context.Param("address")))

	power, err := actions.GetBooted(context, context.Param("address"))
	if err != nil {
		d.Log.Warn("could not get booted status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("received booted status", zap.String("booted", power.Booted))
	context.JSON(http.StatusOK, power)
}

func (d *DeviceManager) setMute(context *gin.Context) {
	d.Log.Debug("setting mute", zap.String("mute", context.Param("mute")), zap.String("address", context.Param("address")))

	mute, err := strconv.ParseBool(context.Param("mute"))
	if err != nil {
		d.Log.Warn("could not set mute. 'mute' parameter not a valid boolean value", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = actions.SetMute(context, context.Param("address"), mute)
	if err != nil {
		d.Log.Warn("failed to set mute", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("successfully set mute", zap.String("mute", strconv.FormatBool(mute)), zap.String("address", context.Param("addres")))
	context.JSON(http.StatusOK, 1)
}

func (d *DeviceManager) getMute(context *gin.Context) {
	d.Log.Debug("getting mute status", zap.String("address", context.Param("address")))

	mute, err := actions.GetMute(context.Param("address"))
	if err != nil {
		d.Log.Warn("failed to get mute status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("received mute status", zap.String("mute", strconv.FormatBool(mute.Muted)))
	context.JSON(http.StatusOK, mute)
}

func (d *DeviceManager) setVolume(context *gin.Context) {
	d.Log.Debug("setting volume", zap.String("volume", context.Param("volume")), zap.String("address", context.Param("address")))

	volume, err := strconv.Atoi(context.Param("volume"))
	if err != nil {
		d.Log.Warn("cannot set volume. 'volume' parameter not a valid integer value", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = actions.SetVolume(context, context.Param("address"), volume)
	if err != nil {
		d.Log.Warn("failed to set volume", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("successfully set volume", zap.String("volume", strconv.Itoa(volume)), zap.String("address", context.Param("address")))
	context.JSON(http.StatusOK, 1)
}

func (d *DeviceManager) getVolume(context *gin.Context) {
	d.Log.Debug("getting volume", zap.String("address", context.Param("address")))

	volume, err := actions.GetVolume(context.Param("address"))
	if err != nil {
		d.Log.Warn("failed to get volume", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("received volume status", zap.String("volume", strconv.Itoa(volume.Volume)), zap.String("address", context.Param("address")))
	context.JSON(http.StatusOK, volume)
}

func (d *DeviceManager) setInput(context *gin.Context) {
	d.Log.Debug("setting input", zap.String("input", context.Param("input")), zap.String("address", context.Param("address")))

	err := actions.SetInput(context, context.Param("address"), context.Param("input"))
	if err != nil {
		d.Log.Warn("failed to set input", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("successfully set input", zap.String("input", context.Param("input")), zap.String("address", context.Param("address")))
	context.JSON(http.StatusOK, 1)
}

func (d *DeviceManager) getInput(context *gin.Context) {
	d.Log.Debug("getting input status", zap.String("address", context.Param("address")))

	input, err := actions.GetInput(context.Param("address"))
	if err != nil {
		d.Log.Warn("failed to get input status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("received input status", zap.String("input", input.Input), zap.String("address", context.Param("address")))
	context.JSON(http.StatusOK, input)
}

func (d *DeviceManager) setBlank(context *gin.Context) {
	d.Log.Debug("setting blank", zap.String("blank", context.Param("blank")), zap.String("address", context.Param("address")))

	blank, err := strconv.ParseBool(context.Param("blank"))
	if err != nil {
		d.Log.Warn("cannot set blank. 'blank' parameter not a valid boolean value", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = actions.SetBlank(context, context.Param("address"), blank)
	if err != nil {
		d.Log.Warn("failed to set blank", zap.Error(err))
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	d.Log.Debug("successfully set blank", zap.String("blank", context.Param("blank")), zap.String("address", context.Param("address")))
	context.JSON(http.StatusOK, 1)
}
