package gateway

import (
	"context"
)

type HelloService struct {
	UnimplementedGreeterServer
}

func (s *HelloService) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: in.Name + " welcome"}, nil
}
