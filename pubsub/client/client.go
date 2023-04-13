package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/go-micro/plugins/v3/broker/redis"
	"github.com/google/uuid"
	pubsub "microDemo/pubsub/proto"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("publisher"),
		micro.Broker(redis.NewBroker()),
		micro.Registry(consul.NewRegistry()),
	)
	pub := micro.NewEvent("test:topic", service.Client())
	t := time.NewTicker(time.Second)
	go func() {
		for _ = range t.C {
			// create new event
			msgId, _ := uuid.NewUUID()
			ev := &pubsub.Event{
				Id:        msgId.String(),
				Timestamp: time.Now().Unix(),
				Message:   fmt.Sprintf("Messaging you all day on %s", "test:topic"),
			}

			fmt.Printf("publishing %+v\n", ev)

			// publish an event
			if err := pub.Publish(context.Background(), ev); err != nil {
				fmt.Printf("error publishing: %v", err)
			}
		}
	}()
	_ = service.Run()
}
