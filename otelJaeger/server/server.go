package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc"
	oteJaeger "microDemo/otelJaeger/proto"
	"net"
	"time"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("api-server"))),
	)
	return tp, nil
}

func main() {
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		fmt.Println(err)
	}
	// 注册我们的 TracerProvider 为全局的 所以任何导入
	// 将来的 instrumentation 将默认使用它
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 创建 子 context  用来传递给 子协程  用于 通信 底层实现的是 channel   cancel 用于关闭用
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 当应用程序退出时，干净地关闭和刷新遥测
	defer func(ctx context.Context) {
		// 在关闭应用程序时，不要使其挂起
		ctx, cancel := context.WithTimeout(ctx, time.Second*5) // 从创建后 超过5秒后 关闭
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil { // Shutdown按注册span处理器的顺序关闭它们
			fmt.Println(err)
		}
	}(ctx)

	server := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	oteJaeger.RegisterGreeterServer(server, &oteJaeger.HelloService{})

	lis, _ := net.Listen("tcp", ":8787")
	_ = server.Serve(lis)
}
