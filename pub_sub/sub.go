package pub_sub

import (
	"go-rabbitmq/base"
)

// Subscribe1 发布/订阅模式订阅者1
func Subscribe1() {
	rabbit := base.NewRabbitMQPublish("NewPub")
	rabbit.ReceivePub()
}
