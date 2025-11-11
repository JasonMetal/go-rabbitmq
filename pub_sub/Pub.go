package pub_sub

import (
	"go-rabbitmq/base"
	"strconv"
	"time"
)

// Publish 发布/订阅模式发布者
func Publish() {
	rabbit := base.NewRabbitMQPublish("NewPub")
	for i := 0; i < 100; i++ {
		rabbit.PublishPub("PubMsg" + strconv.Itoa(i) + "data")
		time.Sleep(1 * time.Second)
	}
}
