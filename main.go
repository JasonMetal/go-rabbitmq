package main

import (
	"fmt"
	"go-rabbitmq/pub_sub"
	"go-rabbitmq/route"
	"go-rabbitmq/simples"
	"go-rabbitmq/topic"
	"go-rabbitmq/worker"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-rabbitmq <mode>")
		fmt.Println("Available modes:")
		fmt.Println("  simple-publish    - Run simple mode publisher")
		fmt.Println("  simple-consume    - Run simple mode consumer")
		fmt.Println("  worker-publish    - Run worker mode publisher")
		fmt.Println("  worker-consume1   - Run worker mode consumer 1")
		fmt.Println("  worker-consume2   - Run worker mode consumer 2")
		fmt.Println("  pub-publish       - Run publish/subscribe mode publisher")
		fmt.Println("  pub-subscribe1    - Run publish/subscribe mode subscriber 1")
		fmt.Println("  pub-subscribe2    - Run publish/subscribe mode subscriber 2")
		fmt.Println("  route-publish     - Run routing mode publisher")
		fmt.Println("  route-consume-one - Run routing mode consumer for 'one' key")
		fmt.Println("  route-consume-two - Run routing mode consumer for 'two' key")
		fmt.Println("  topic-publish     - Run topic mode publisher")
		fmt.Println("  topic-consume     - Run topic mode consumer")
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
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
	}
}
