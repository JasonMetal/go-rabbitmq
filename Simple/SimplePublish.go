package Worker

import (
	"fmt"
	"go-rabbitmq/Base"
	"strconv"
	"time"
)

//worker模式下
func main() {
	rabbitmq := Base.NewRabbitMQSimple("Simple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("hello worker mode" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
