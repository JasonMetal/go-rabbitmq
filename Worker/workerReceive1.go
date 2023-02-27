package Worker

import (
	"go-rabbitmq/Base"
)

//worker模式下
func main() {
	rabbitmq := Base.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumeSimple()
}
