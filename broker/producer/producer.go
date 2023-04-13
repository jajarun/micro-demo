package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asim/go-micro/v3/broker"
	"github.com/go-micro/plugins/v3/broker/redis"
)

var (
	topic = "go.micro.topic.foo"
)

func pub(redisBroker broker.Broker) {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := redisBroker.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func main() {
	redisBroker := redis.NewBroker(
		broker.Addrs("127.0.0.1:6379"),
	)
	if err := redisBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := redisBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	pub(redisBroker)
}
