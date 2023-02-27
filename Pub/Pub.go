package Pub

import (
	"go-rabbitmq/Base"
	"strconv"
	"time"
)

func main() {
	rabbit := Base.NewRabbitMQPublish("NewPub")
	for i := 0; i < 100; i++ {
		rabbit.PublishPub("PubMsg" + strconv.Itoa(i) + "data")
		time.Sleep(1 * time.Second)
	}
}
