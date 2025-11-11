package worker

import (
	"fmt"
	"go-rabbitmq/base"
	"strconv"
	"time"
)

// WorkerPublish 工作队列模式发布者
func WorkerPublish() {
	rabbitmq := base.NewRabbitMQSimple("Simple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("hello worker mode" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
