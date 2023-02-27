package Route

import (
	"fmt"
	"go-rabbitmq/Base"
	"strconv"
	"time"
)

func main() {
	one := Base.NewRabbitMQRouting("exchangeRoute", "one")
	two := Base.NewRabbitMQRouting("exchangeRoute", "two")
	for i := 0; i < 100; i++ {
		one.PublishRouting("hello one" + strconv.Itoa(i))
		two.PublishRouting("hello two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("num is ", i)
	}
}
