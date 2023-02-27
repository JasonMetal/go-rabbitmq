package Route

import "go-rabbitmq/Base"

func main()  {
	rabbit := Base.NewRabbitMQRouting("exchangeSth","one")
	rabbit.ReceiveRouting()
}
