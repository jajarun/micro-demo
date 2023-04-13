package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/go-micro/metadata"
	pubsub "microDemo/pubsub/proto"
)

type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *pubsub.Event) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	servive := micro.NewService(
		micro.Name("subscribe"),
		//micro.Broker(redis.NewBroker()),
		micro.Registry(consul.NewRegistry()),
	)
	//_ = micro.RegisterSubscriber("test:topic", servive.Server(), new(Sub))
	_ = micro.RegisterSubscriber("test:topic", servive.Server(), new(Sub), server.SubscriberQueue("queue1"))
	_ = servive.Run()
}
