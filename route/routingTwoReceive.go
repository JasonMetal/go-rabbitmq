package route

import "go-rabbitmq/base"

func RouteConsumeTwo() {
	rabbit := base.NewRabbitMQRouting("exchangeSth", "two")
	rabbit.ReceiveRouting()
}
