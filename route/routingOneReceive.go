package route

import "go-rabbitmq/base"

func RouteConsumeOne() {
	rabbit := base.NewRabbitMQRouting("exchangeSth", "one")
	rabbit.ReceiveRouting()
}
