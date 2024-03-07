package main

import (
	"05_sever_unary_interceptors/src/pb/hello"
	"context"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestServerUnaryInterceptors(t *testing.T) {
	conn, err := grpc.Dial("localhost:4926", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancelFunc()

	if _, err := client.Hello(ctx, &hello.HelloRequest{}); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
