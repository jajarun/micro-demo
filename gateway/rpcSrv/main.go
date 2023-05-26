package main

import (
	"google.golang.org/grpc"
	gateway "microDemo/gateway/proto"
	"net"
)

func main() {

	server := grpc.NewServer()

	gateway.RegisterGreeterServer(server, &gateway.HelloService{})

	lis, _ := net.Listen("tcp", ":9090")
	_ = server.Serve(lis)
}
