package device

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceManager struct {
	Log *zap.Logger
}

func (d *DeviceManager) RunHTTPServer(router *gin.Engine, port string) error {
	d.Log.Info("registering http endpoints")

	// action endpoints
	route := router.Group("/api/v1")
	route.GET("/:address/power/:power", d.setPower)
	route.GET("/:address/volume/:volume", d.setVolume)
	route.GET("/:address/volume/mute/:mute", d.setMute)
	route.GET("/:address/display/:blank", d.setBlank)
	route.GET("/:address/input/:input", d.setInput)

	// status endpoints
	route.GET("/:address/power", d.getPower)
	route.GET("/:address/volume", d.getVolume)
	route.GET("/:address/volume/mute", d.getMute)
	route.GET("/:address/input", d.getInput)
	route.GET("/:address/booted", d.getBooted)

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	d.Log.Info("running http server")
	router.Run(server.Addr)

	return fmt.Errorf("http server stopped")
}
