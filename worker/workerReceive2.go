package worker

import (
	"go-rabbitmq/base"
)

// WorkerReceive2 工作队列模式消费者2
func WorkerReceive2() {
	rabbitmq := base.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumeSimple()
}
