package main

import (
	"client_stream_rpc/src/pb/hello"
	"client_stream_rpc/src/serverImp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	listen, err := net.Listen("tcp", "localhost:4926")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &serverImp.HelloServer{})
	defer server.Stop()
	server.Serve(listen)
}
