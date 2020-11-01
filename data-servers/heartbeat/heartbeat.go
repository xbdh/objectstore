package heartbeat

import (
	"data-servers/config"
	"data-servers/rabbitmq"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(config.RabbitmqServe)
	defer q.Close()
	for {
		q.Publish("apiServers", config.ListenAddress)
		time.Sleep(5 * time.Second)
	}
}