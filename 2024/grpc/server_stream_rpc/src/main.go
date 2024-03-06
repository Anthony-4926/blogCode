package main

import (
	"google.golang.org/grpc"
	"net"
	"server_stream_rpc/src/pb/hello"
	"server_stream_rpc/src/serverImp"
)

func main() {
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
