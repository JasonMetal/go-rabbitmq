package Topic

import "go-rabbitmq/Base"

func main() {
	rabbit := Base.NewRabbitMQTopic("exchangeTopic", "#")
	rabbit.RecieveTopic()
}
