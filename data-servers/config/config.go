package config

import "flag"

//const (
//	RabbitmqServe = "amqp://xbdh:0315@172.17.0.3:5672/object-storage"
//
//	ListenAdress ="10.29.2.1:12345"
//)
var (
	RabbitmqServe string
	//= "amqp://xbdh:0315@172.17.0.3:5672/object-storage"
	//StorageRoot= "/home/chen/file/storage"
	ListenAddress string
	//= "10.29.1.1:12345"
	StorageRoot string
	//"/home/chen/file/storage/"
)

func Init()  {
	flag.StringVar(&RabbitmqServe ,"RabbitmqServe","amqp://xbdh:0315@172.17.0.3:5672/object-storage","消息队列地址")
	flag.StringVar(&ListenAddress ,"ListenAddress","","监听地址")
	flag.StringVar(&StorageRoot ,"StorageRoot","","存储根地址")
	//一定要解析啊
	flag.Parse()
}

