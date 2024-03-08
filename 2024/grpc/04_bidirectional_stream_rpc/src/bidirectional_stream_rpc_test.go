package main

import (
	"04_bidirectional_stream_rpc/src/pb/hello"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"testing"
	"time"
)

func TestBidirectionalStreamRpc(t *testing.T) {
	conn, err := grpc.Dial("localhost:4926", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancelFunc()

	stream, err := client.Hello(ctx)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("hello world %d", i)
			if err := stream.Send(&hello.HelloRequest{Msg: msg}); err != nil {
				logrus.Warnln(err)
			}
		}
		stream.CloseSend()
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			logrus.Infoln(err)
			break
		}
		logrus.Infoln(resp)
	}

	conn.Close()

}
