package main

import (
	"04_bidirectional_stream_rpc/src/pb/hello"
	imp "04_bidirectional_stream_rpc/src/severImp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	listener, err := net.Listen("tcp", "localhost:4926")
	if err != nil {
		logrus.Fatal(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	hello.RegisterHelloServiceServer(grpcServer, &imp.HelloSeverImp{})

	if err = grpcServer.Serve(listener); err != nil {
		logrus.Fatal(err)
	}

}
