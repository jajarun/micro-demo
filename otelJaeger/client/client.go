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
	"google.golang.org/grpc/credentials/insecure"
	oteJaeger "microDemo/otelJaeger/proto"
	"time"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("api-client"))),
	)
	return tp, nil
}

func main() {
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		fmt.Println(err)
		return
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

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
	//tr := otel.Tracer("component-client")
	//_, span := tr.Start(ctx, "startCall")
	conn, _ := grpc.Dial(
		"10.1.3.105:8787",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),   // 拦截普通的一次请求一次响应的rpc服务
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()), // 拦截流式的rpc服务
	)
	defer conn.Close()
	client := oteJaeger.NewGreeterClient(conn)
	reply, _ := client.SayHello(context.Background(), &oteJaeger.HelloRequest{Name: "jaja"})
	fmt.Println(reply.Message)
	//defer span.End()
}
