package transaction

import (
	"fmt"
	"go-rabbitmq/base"
	"log"
)

// TransactionConsumeSimple 简单模式事务消费者示例
func TransactionConsumeSimple() {
	rabbitmq := base.NewRabbitMQSimple("SimpleTx")
	defer rabbitmq.Destory()

	fmt.Println("等待接收简单模式事务消息...")
	rabbitmq.ConsumeSimple()
}

// TransactionConsumePub 发布/订阅模式事务消费者示例
func TransactionConsumePub() {
	rabbitmq := base.NewRabbitMQPublish("NewPubTx")
	defer rabbitmq.Destory()

	fmt.Println("等待接收发布/订阅模式事务消息...")
	rabbitmq.ReceivePub()
}

// TransactionConsumeRoutingOne 路由模式事务消费者示例 (接收 "one" 路由键的消息)
func TransactionConsumeRoutingOne() {
	rabbitmq := base.NewRabbitMQRouting("exchangeRouteTx", "one")
	defer rabbitmq.Destory()

	fmt.Println("等待接收路由模式事务消息 (路由键: one)...")
	rabbitmq.ReceiveRouting()
}

// TransactionConsumeRoutingTwo 路由模式事务消费者示例 (接收 "two" 路由键的消息)
func TransactionConsumeRoutingTwo() {
	rabbitmq := base.NewRabbitMQRouting("exchangeRouteTx", "two")
	defer rabbitmq.Destory()

	fmt.Println("等待接收路由模式事务消息 (路由键: two)...")
	rabbitmq.ReceiveRouting()
}

// TransactionConsumeTopic 话题模式事务消费者示例 (接收所有话题消息)
func TransactionConsumeTopic() {
	rabbitmq := base.NewRabbitMQTopic("exchangeTopicTx", "#")
	defer rabbitmq.Destory()

	fmt.Println("等待接收话题模式事务消息...")
	rabbitmq.RecieveTopic()
}

// SimulateTransactionFailure 模拟事务失败的示例
func SimulateTransactionFailure() {
	rabbitmq := base.NewRabbitMQSimple("SimpleTx")
	defer rabbitmq.Destory()

	// 开启事务
	err := rabbitmq.TxSelect()
	if err != nil {
		log.Fatalf("开启事务失败: %s", err)
	}

	// 声明队列
	_, err = rabbitmq.Channel.QueueDeclare(
		rabbitmq.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = rabbitmq.TxRollback()
		log.Fatalf("队列声明失败并已回滚: %s", err)
		return
	}

	// 模拟一个错误情况（例如网络问题导致无法发布消息）
	// 在实际应用中，这可能是由于某些业务逻辑检查失败等原因
	simulatedError := true
	if simulatedError {
		_ = rabbitmq.TxRollback()
		fmt.Println("模拟事务回滚成功")
		return
	}

	// 正常流程会提交事务
	/*
		err = rabbitmq.TxCommit()
		if err != nil {
			_ = rabbitmq.TxRollback()
			log.Fatalf("事务提交失败并已回滚: %s", err)
			return
		}
	*/
}
