package oteJaeger

import (
	"context"
	"time"
)

type HelloService struct {
	UnimplementedGreeterServer
}

func (s *HelloService) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	//tr := otel.Tracer("component-server")
	//_, span := tr.Start(ctx, "SayHello")
	time.Sleep(time.Second)
	//defer span.End()
	return &HelloReply{Message: in.Name + " welcome"}, nil
}
