package main

import (
	"fmt"
	"github.com/go-micro/plugins/v3/broker/redis"
	"log"

	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/cmd"
)

var (
	topic = "go.micro.topic.foo"
)

// Example of a shared subscription which receives a subset of messages
//func sharedSub() {
//	_, err := broker.Subscribe(topic, func(p broker.Event) error {
//		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
//		return nil
//	}, broker.Queue("consumer"))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

// Example of a subscription which receives all the messages
func sub(redisBroker broker.Broker) {
	_, err := redisBroker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	redisBroker := redis.NewBroker()
	if err := redisBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := redisBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	sub(redisBroker)
	//sharedSub()
	select {}
}
