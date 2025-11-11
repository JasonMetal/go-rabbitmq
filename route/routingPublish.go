package route

import (
	"fmt"
	"go-rabbitmq/base"
	"strconv"
	"time"
)

func RoutePublish() {
	one := base.NewRabbitMQRouting("exchangeRoute", "one")
	two := base.NewRabbitMQRouting("exchangeRoute", "two")
	for i := 0; i < 100; i++ {
		one.PublishRouting("hello one" + strconv.Itoa(i))
		two.PublishRouting("hello two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("num is ", i)
	}
}
