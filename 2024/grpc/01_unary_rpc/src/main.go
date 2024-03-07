package main

import (
	"google.golang.org/grpc"
	"net"
	Hello "pprof/src/pb/this_is_go_package_content"
	imp "pprof/src/serverImp"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:4096")

	if err != nil {
		panic(err)
	}
	defer listen.Close()

	server := grpc.NewServer()
	defer server.Stop()
	Hello.RegisterHelloServiceServer(server, &imp.HelloServer{})

	if err := server.Serve(listen); err != nil {
		return
	}
}
