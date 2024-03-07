package main

import (
	"05_sever_unary_interceptors/src/pb/hello"
	"05_sever_unary_interceptors/src/serverImp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	listener, err := net.Listen("tcp", "localhost:4926")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(serverImp.TimeCounsumeMiddleware()), // 注册中间件
	)
	defer server.Stop()

	hello.RegisterHelloServiceServer(server, &serverImp.HelloServer{})

	if err = server.Serve(listener); err != nil {
		panic(err)
	}

}
