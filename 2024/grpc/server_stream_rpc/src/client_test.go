package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"server_stream_rpc/src/pb/hello"
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

	stream, err := client.Hello(ctx, &hello.HelloRequest{Who: &hello.People{Name: "Anthony"}})
	if err != nil {
		t.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Warning(err)
		} else {
			fmt.Println(resp)
		}
	}
}
