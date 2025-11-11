package topic

import (
	"fmt"
	"go-rabbitmq/base"
	"strconv"
	"time"
)

func TopicPublish() {
	three := base.NewRabbitMQTopic("exchangeTopic", "topic11.three")
	four := base.NewRabbitMQTopic("exchangeTopic", "topic11.four")

	for i := 0; i <= 10; i++ {
		three.PublishTopic("hello topic three" + strconv.Itoa(i))
		four.PublishTopic("hello topic four" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("num is :", i)
	}

}
