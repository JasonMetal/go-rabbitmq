package simples

import (
	"fmt"
	"go-rabbitmq/base"
	"strconv"
	"time"
)

// SimplePublish 简单模式发布者
func SimplePublish() {
	rabbitmq := base.NewRabbitMQSimple("Simple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("hello worker mode" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
