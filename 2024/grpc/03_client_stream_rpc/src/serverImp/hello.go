package serverImp

import (
	"client_stream_rpc/src/pb/hello"
	"github.com/sirupsen/logrus"
	"io"
)

type HelloServer struct {
	hello.UnimplementedHelloServiceServer
}

func (h HelloServer) Hello(server hello.HelloService_HelloServer) error {
	for {
		request, err := server.Recv()
		if err == io.EOF {
			server.SendAndClose(&hello.HelloResponse{Msg: "Tom"})
			break
		}
		if err != nil {
			return err
		} else {
			logrus.Info("recv msg: ", request)
		}
	}
	return nil
}
