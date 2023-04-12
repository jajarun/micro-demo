package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	handler "microDemo/rpcServer/handle"
	rpcServer "microDemo/rpcServer/proto"
)

func main() {
	srv := micro.NewService(
		micro.Name("micro-rpc"),
		micro.Registry(consul.NewRegistry()),
	)
	_ = rpcServer.RegisterGreeterHandler(srv.Server(), new(handler.Greeter))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
