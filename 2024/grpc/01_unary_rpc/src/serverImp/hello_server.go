package serverImp

import (
	"context"
	Hello "pprof/src/pb/this_is_go_package_content"
)

type HelloServer struct {
	Hello.UnimplementedHelloServiceServer
}

func (h HelloServer) Hello(ctx context.Context, request *Hello.HelloRequest) (*Hello.HelloResponse, error) {
	resMsg := &Hello.HelloResponse{
		Who: &Hello.People{
			Name: "Tom",
		},
		Msg: "Hello " + request.Who.Name,
	}
	return resMsg, nil
}
