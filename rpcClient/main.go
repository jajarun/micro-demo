package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	rpcServer "microDemo/rpcServer/proto"
)

func main() {
	//reg := consul.NewRegistry()
	//srvs, _ := reg.GetService("micro-rpc")
	//for _, v := range srvs {
	//	fmt.Println(v.Nodes[0])
	//	fmt.Println(v.Nodes[1])
	//}

	client := rpcServer.NewGreeterService("micro-rpc", micro.NewService(
		micro.Registry(consul.NewRegistry()),
	).Client())
	reply, err := client.SayHello(context.Background(), &rpcServer.HelloRequest{
		Name: "jajarun",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply.Message)
	}
	reply, err = client.SayHelloAgain(context.Background(), &rpcServer.HelloRequest{
		Name: "jajarun",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply.Message)
	}
}
