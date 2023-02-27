package Route

import "go-rabbitmq/Base"

func main() {
	rabbit := Base.NewRabbitMQRouting("exchangeSth", "two")
	rabbit.ReceiveRouting()
}
