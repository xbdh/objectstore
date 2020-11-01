package config

import "flag"

var (
	RabbitmqServe string
	//= "amqp://xbdh:0315@172.17.0.3:5672/object-storage"
	//StorageRoot= "/home/chen/file/storage"
	ListenAddress string

	//= "10.29.1.1:12345"

	ElasticsearchAddress string
)

func Init() {
	flag.StringVar(&RabbitmqServe, "RabbitmqServe", "amqp://xbdh:0315@172.17.0.3:5672/object-storage", "消息队列地址")
	flag.StringVar(&ListenAddress, "ListenAddress", "", "rabbitmq监听地址")
	flag.StringVar(&ElasticsearchAddress, "ElasticsearchAddress", "", "elasticsearch监听地址")
	flag.Parse()
}
