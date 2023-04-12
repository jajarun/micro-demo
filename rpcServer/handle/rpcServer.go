package handler

import (
	"context"
	"fmt"
	rpcServer "microDemo/rpcServer/proto"
)

type Greeter struct{}

func (c *Greeter) SayHello(ctx context.Context, req *rpcServer.HelloRequest, res *rpcServer.HelloReply) error {
	fmt.Println("Received MicroTest.Call request")
	res.Message = "Hello " + req.Name
	return nil
}

func (c *Greeter) SayHelloAgain(ctx context.Context, req *rpcServer.HelloRequest, res *rpcServer.HelloReply) error {
	fmt.Println("Received MicroTest.Call request agin")
	res.Message = "Hello again" + req.Name
	return nil
}
