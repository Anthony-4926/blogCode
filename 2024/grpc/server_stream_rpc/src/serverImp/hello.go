package serverImp

import (
	"fmt"
	"server_stream_rpc/src/pb/hello"
)

type HelloServer struct {
	hello.UnimplementedHelloServiceServer
}

func (h HelloServer) Hello(request *hello.HelloRequest, server hello.HelloService_HelloServer) error {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("hello %d", i)
		server.Send(&hello.HelloResponse{Who: &hello.People{Name: "Tom", Age: 10}, Msg: msg})
	}
	return nil
}
