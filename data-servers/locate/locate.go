package locate

import (
	"data-servers/config"
	"data-servers/rabbitmq"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(config.RabbitmqServe)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		name, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(config.StorageRoot+name) {
			q.Send(msg.ReplyTo, config.ListenAddress)
		}
	}
}
