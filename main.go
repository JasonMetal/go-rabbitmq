package main

import (
	"fmt"
	"go-rabbitmq/pub_sub"
	"go-rabbitmq/route"
	"go-rabbitmq/simples"
	"go-rabbitmq/topic"
	"go-rabbitmq/transaction"
	"go-rabbitmq/worker"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-rabbitmq <mode>")
		fmt.Println("Available modes:")
		fmt.Println("  simple-publish          - Run simple mode publisher")
		fmt.Println("  simple-consume          - Run simple mode consumer")
		fmt.Println("  worker-publish          - Run worker mode publisher")
		fmt.Println("  worker-consume1         - Run worker mode consumer 1")
		fmt.Println("  worker-consume2         - Run worker mode consumer 2")
		fmt.Println("  pub-publish             - Run publish/subscribe mode publisher")
		fmt.Println("  pub-subscribe1          - Run publish/subscribe mode subscriber 1")
		fmt.Println("  pub-subscribe2          - Run publish/subscribe mode subscriber 2")
		fmt.Println("  route-publish           - Run routing mode publisher")
		fmt.Println("  route-consume-one       - Run routing mode consumer for 'one' key")
		fmt.Println("  route-consume-two       - Run routing mode consumer for 'two' key")
		fmt.Println("  topic-publish           - Run topic mode publisher")
		fmt.Println("  topic-consume           - Run topic mode consumer")
		fmt.Println("  tx-simple-publish       - Run transactional simple mode publisher")
		fmt.Println("  tx-simple-consume       - Run transactional simple mode consumer")
		fmt.Println("  tx-pub-publish          - Run transactional publish/subscribe mode publisher")
		fmt.Println("  tx-pub-subscribe        - Run transactional publish/subscribe mode subscriber")
		fmt.Println("  tx-route-publish        - Run transactional routing mode publisher")
		fmt.Println("  tx-route-consume-one    - Run transactional routing mode consumer for 'one' key")
		fmt.Println("  tx-route-consume-two    - Run transactional routing mode consumer for 'two' key")
		fmt.Println("  tx-topic-publish        - Run transactional topic mode publisher")
		fmt.Println("  tx-topic-consume        - Run transactional topic mode consumer")
		fmt.Println("  tx-simulate-failure     - Simulate transaction failure")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "simple-publish":
		simples.SimplePublish()
	case "simple-consume":
		simples.SimpleReceive()
	case "worker-publish":
		worker.WorkerPublish()
	case "worker-consume1":
		worker.WorkerReceive1()
	case "worker-consume2":
		worker.WorkerReceive2()
	case "pub-publish":
		pub_sub.Publish()
	case "pub-subscribe1":
		pub_sub.Subscribe1()
	case "pub-subscribe2":
		pub_sub.Subscribe2()
	case "route-publish":
		route.RoutePublish()
	case "route-consume-one":
		route.RouteConsumeOne()
	case "route-consume-two":
		route.RouteConsumeTwo()
	case "topic-publish":
		topic.TopicPublish()
	case "topic-consume":
		topic.TopicReceive()
	case "tx-simple-publish":
		transaction.TransactionPublishSimple()
	case "tx-simple-consume":
		transaction.TransactionConsumeSimple()
	case "tx-pub-publish":
		transaction.TransactionPublishPub()
	case "tx-pub-subscribe":
		transaction.TransactionConsumePub()
	case "tx-route-publish":
		transaction.TransactionPublishRouting()
	case "tx-route-consume-one":
		transaction.TransactionConsumeRoutingOne()
	case "tx-route-consume-two":
		transaction.TransactionConsumeRoutingTwo()
	case "tx-topic-publish":
		transaction.TransactionPublishTopic()
	case "tx-topic-consume":
		transaction.TransactionConsumeTopic()
	case "tx-simulate-failure":
		transaction.SimulateTransactionFailure()
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
	}
}
