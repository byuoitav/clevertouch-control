package main

import (
	"net/http"

	"github.com/byuoitav/clevertouch-control/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8007"

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// action endpoints
	write := router.Group("/write")
	write.GET("/:address/power/on", handlers.PowerOn)
	write.GET("/:address/power/standby", handlers.Standby)
	write.GET("/:address/input/:port", handlers.SwitchInput)
	write.GET("/:address/volume/set/:value", handlers.SetVolume)
	write.GET("/:address/volume/mute", handlers.VolumeMute)
	write.GET("/:address/volume/unmute", handlers.VolumeUnmute)
	write.GET("/:address/display/blank", handlers.BlankDisplay)
	write.GET("/:address/display/unblank", handlers.UnblankDisplay)

	// status endpoints
	read := router.Group("/read")
	read.GET("/:address/input/current")
	read.GET("/:address/power/status", handlers.GetPower)
	read.GET("/:address/input/current", handlers.GetInput)
	read.GET("/:address/input/list", handlers.GetInputList)
	read.GET("/:address/active/:port", handlers.GetActiveSignal)
	read.GET("/:address/volume/level", handlers.GetVolume)
	read.GET("/:address/volume/mute/status", handlers.GetMute)
	read.GET("/:address/display/status", handlers.GetBlank)
	read.GET("/:address/hardware", handlers.GetHardwareInfo)

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	router.Run(server.Addr)
}
