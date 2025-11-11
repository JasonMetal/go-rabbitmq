package topic

import "go-rabbitmq/base"

func TopicReceive() {
	rabbit := base.NewRabbitMQTopic("exchangeTopic", "#")
	rabbit.RecieveTopic()
}
