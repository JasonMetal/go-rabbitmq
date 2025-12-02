package transaction

import (
	"fmt"
	"go-rabbitmq/base"
	"log"
	"strconv"
	"time"
)

// TransactionPublishSimple 简单模式事务发布示例
func TransactionPublishSimple() {
	rabbitmq := base.NewRabbitMQSimple("SimpleTx")

	//for i := 0; i <= 10; i++ {
	for i := 0; ; i++ {
		msg := "hello transaction simple mode " + strconv.Itoa(i)
		err := rabbitmq.PublishSimpleWithTx(msg)
		if err != nil {
			log.Printf("发送消息失败: %s\n", err)
		} else {
			fmt.Printf("成功发送消息: %s\n", msg)
		}
		time.Sleep(1 * time.Second)
	}
}

// TransactionPublishPub 发布/订阅模式事务发布示例
func TransactionPublishPub() {
	rabbitmq := base.NewRabbitMQPublish("NewPubTx")

	for i := 0; ; i++ {
		msg := "hello transaction pub/sub mode " + strconv.Itoa(i)
		err := rabbitmq.PublishPubWithTx(msg)
		if err != nil {
			log.Printf("发送消息失败: %s\n", err)
		} else {
			fmt.Printf("成功发送消息: %s\n", msg)
		}
		time.Sleep(1 * time.Second)
	}
}

// TransactionPublishRouting 路由模式事务发布示例
func TransactionPublishRouting() {
	one := base.NewRabbitMQRouting("exchangeRouteTx", "one")
	two := base.NewRabbitMQRouting("exchangeRouteTx", "two")

	for i := 0; i <= 5; i++ {
		msgOne := "hello transaction routing one " + strconv.Itoa(i)
		err := one.PublishRoutingWithTx(msgOne)
		if err != nil {
			log.Printf("发送消息失败: %s\n", err)
		} else {
			fmt.Printf("成功发送消息: %s\n", msgOne)
		}

		msgTwo := "hello transaction routing two " + strconv.Itoa(i)
		err = two.PublishRoutingWithTx(msgTwo)
		if err != nil {
			log.Printf("发送消息失败: %s\n", err)
		} else {
			fmt.Printf("成功发送消息: %s\n", msgTwo)
		}

		time.Sleep(1 * time.Second)
	}
}

// TransactionPublishTopic 话题模式事务发布示例
func TransactionPublishTopic() {
	three := base.NewRabbitMQTopic("exchangeTopicTx", "topic.three")
	four := base.NewRabbitMQTopic("exchangeTopicTx", "topic.four")

	for i := 0; i <= 5; i++ {
		msgThree := "hello transaction topic three " + strconv.Itoa(i)
		err := three.PublishTopicWithTx(msgThree)
		if err != nil {
			log.Printf("发送消息失败: %s\n", err)
		} else {
			fmt.Printf("成功发送消息: %s\n", msgThree)
		}

		msgFour := "hello transaction topic four " + strconv.Itoa(i)
		err = four.PublishTopicWithTx(msgFour)
		if err != nil {
			log.Printf("发送消息失败: %s\n", err)
		} else {
			fmt.Printf("成功发送消息: %s\n", msgFour)
		}

		time.Sleep(1 * time.Second)
	}
}
