package Base

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//连接
const URL = "amqp://guest:guest@localhost:5672/"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     URL,
	}
}

//断开连接（channel，connection），销毁方法
func (r RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~简单模式 start

// NewRabbitMQSimple
//  @Description:
//  @param queueName
//  @return *RabbitMQ
//
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接失败！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建channel失败")
	return rabbitmq
}

//
// PublishSimple
//  @Description: 生产者-简单模式
//  @receiver r
//  @param msg
//
func (r *RabbitMQ) PublishSimple(msg string) {
	_, err := r.channel.QueueDeclare(
		//name string, durable, autoDelete, exclusive, noWait bool, args Table
		r.QueueName, //队列名，
		false,       //durable，是否持久化
		false,       //autoDelete，自动删除
		false,       //exclusive，是否排他
		false,       //noWait，是否阻塞
		nil,         //args，额外参数属性
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("PublishSimple 's QueueDeclare err is :", err)
	}

	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		//发送的消息
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg), //消息内容
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	_, err := r.channel.QueueDeclare(
		r.QueueName, //队列名，
		false,       //durable，是否持久化
		false,       //autoDelete，自动删除
		false,       //exclusive，是否排他
		false,       //noWait，是否阻塞
		nil,         //args，额外参数属性
	)
	if err != nil {
		fmt.Println("ConsumeSimple 's QueueDeclare err is :", err)
	}
	msgs, errs := r.channel.Consume(
		r.QueueName,
		"",    //区分多个消费者
		true,  // 是否自动应答
		false, //
		false, // true：不能及将同一个connection中发送的消息，传递这个connection中的消费者
		false,
		nil,
	)
	if errs != nil {
		fmt.Println("ConsumeSimple 's Consume err is :", errs)
	}

	//消费队列
	forever := make(chan bool)
	go func() {
		//消费队列的逻辑
		//入库固化，及其他操作
		for msg := range msgs {
			log.Printf("consume one msg :%s", msg.Body)
		}
	}()
	log.Printf("退出ctrl+c")
	<-forever
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~简单模式 end

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~订阅模式 start

func NewRabbitMQPublish(exchangeName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接失败")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建channel失败")
	return rabbitmq
}

func (r *RabbitMQ) PublishPub(msg string) {
	//1. 创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	//2. 发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func (r *RabbitMQ) ReceivePub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	//2. 尝试创建一个队列
	declare, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//3. 绑定队列到exchange
	err = r.channel.QueueBind(
		declare.Name,
		"",
		r.Exchange, //交换机名字
		false,
		nil,
	)
	//消费者
	consume, errConsume := r.channel.Consume(
		declare.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if errConsume != nil {
		fmt.Println("ReceivePub 's errConsume is :", errConsume)
	}
	forever := make(chan bool)
	go func() {
		for msg := range consume {
			log.Printf("消费一个消息 ：%s", msg.Body)
		}
	}()
	log.Printf("退出，ctrl+c")
	<-forever //等待消息阻塞
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~订阅模式 end

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~路由模式 start
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接失败！")
	//获取Channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建channel失败")
	return rabbitmq
}

//路由模式发送消息
func (r *RabbitMQ) PublishRouting(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct", //点对点
		true,
		false, false, false, nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

//第三步： routing模式下的消息消费
func (r *RabbitMQ) ReceiveRouting() {
	//1.
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true, //是否持久化
		false,
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定，一般设置为false
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	//声明队列
	//2.尝试创建一个队列，这里注意队列名为空
	declare, err := r.channel.QueueDeclare(
		"", //队列名为空
		false,
		false,
		true, //排他性
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//3.绑定队列到exchange中
	err = r.channel.QueueBind(
		declare.Name,
		r.Key,      //pub/sub下，key要为空
		r.Exchange, //交换机名
		false,
		nil,
	)
	//消费消息
	consume, err := r.channel.Consume(
		declare.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	forever := make(chan bool)
	go func() {
		for msg := range consume {
			log.Printf("Received a msg :%s", msg.Body)
		}
	}()
	log.Printf("退出ctrl+c")
	<-forever
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~路由模式 end

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~话题模式 start
func NewRabbitMQTopic(exchange string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchange, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "Failed to connect rabbitmq!")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "Failed to open a channel")
	return rabbitmq
}

func (r *RabbitMQ) PublishTopic(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	//2.生产消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key, //要设置
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func (r *RabbitMQ) RecieveTopic() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange!")
	//2.尝试创建队列
	declare, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//3.绑定队列到exchange中
	err = r.channel.QueueBind(
		declare.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	//消费消息
	msgs, err := r.channel.Consume(
		declare.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a msg :%s", msg.Body)
		}
	}()
	fmt.Println("退出ctrl+c")
	<-forever
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~话题模式 end

//
