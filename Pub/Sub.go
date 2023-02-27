package Pub

import (
	"go-rabbitmq/Base"
)

func main() {
	rabbit := Base.NewRabbitMQPublish("NewPub")
	rabbit.ReceivePub()
}
