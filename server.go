package main

import (
	"net/http"

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
	write.GET("/:address/power/standby")

	// status endpoints
	read := router.Group("/read")
	read.GET("/:address/input/current")

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	router.Run(server.Addr)
}
