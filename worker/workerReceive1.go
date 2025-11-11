package worker

import (
	"go-rabbitmq/base"
)

// WorkerReceive1 工作队列模式消费者1
func WorkerReceive1() {
	rabbitmq := base.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumeSimple()
}
