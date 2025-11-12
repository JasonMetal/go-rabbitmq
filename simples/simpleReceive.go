package simples

import (
	"go-rabbitmq/base"
)

// SimpleReceive 简单模式消费者
func SimpleReceive() {
	rabbitmq := base.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumeSimple()
}
