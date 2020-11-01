package main

import (
	"api-servers/config"
	"api-servers/heartbeat"
	"api-servers/locate"
	"api-servers/objects"
	"api-servers/version"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	go heartbeat.ListenHeartbeat()

	r := gin.Default()
	r.GET("/object/:name", objects.Get)
	r.POST("/object/:name", objects.Post)

	r.GET("/locate/:name", locate.Get)

	r.GET("/version/:name", version.Get)
	r.Run(config.ListenAddress)
}
