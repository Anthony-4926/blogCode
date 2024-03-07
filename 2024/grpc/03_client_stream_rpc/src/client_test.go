package main

import (
	"client_stream_rpc/src/pb/hello"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestServerStreamRpc(t *testing.T) {
	conn, err := grpc.Dial("localhost:4926", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	client := hello.NewHelloServiceClient(conn)

	stream, err := client.Hello(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		stream.Send(&hello.HelloRequest{Who: &hello.People{Name: "Anthony", Age: uint32(i)}})
	}
	recv, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatal(err)
	}
	logrus.Info(recv.Msg)
}
