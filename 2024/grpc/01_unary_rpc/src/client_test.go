package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	Hello "pprof/src/pb/this_is_go_package_content"
	"testing"
)

func TestUnaryRpc(t *testing.T) {
	conn, err := grpc.Dial("localhost:4096", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := Hello.NewHelloServiceClient(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	req := &Hello.HelloRequest{
		Who: &Hello.People{
			Name: "Anthony",
		},
		Msg: "Hello",
	}
	resp, err := client.Hello(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Print(resp)
}
