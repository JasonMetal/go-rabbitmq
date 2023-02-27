package Worker

import (
	"go-rabbitmq/Base"
)

//worker模式下 多一个消费端
func main() {
	rabbitmq := Base.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumeSimple()
}
