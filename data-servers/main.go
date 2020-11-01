package main

import (
	"data-servers/config"
	"data-servers/heartbeat"
	"data-servers/locate"
	"data-servers/objects"
	"github.com/gin-gonic/gin"
)



func main()  {
	config.Init()
	//time.Sleep(time.Second*2)
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()

	r:=gin.Default()
	r.GET("/object/:name",objects.Get)
	r.POST("/object/:name",objects.Post)
	r.Run(config.ListenAddress)
}