package locate

import (
	"api-servers/config"
	"api-servers/rabbitmq"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//const RabbitmqServe = "amqp://xbdh:0315@locahost:5673"

func Locate(name string) string {
	q := rabbitmq.New(config.RabbitmqServe)
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}

func Get(c *gin.Context) {
	name := c.Param("name")
	info := Locate(name)
	if len(info) == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	c.Writer.Write(b)
}
