package pub_sub

import (
	"go-rabbitmq/base"
)

// Subscribe2 发布/订阅模式订阅者2
func Subscribe2() {
	rabbit := base.NewRabbitMQPublish("NewPub")
	rabbit.ReceivePub()
}
