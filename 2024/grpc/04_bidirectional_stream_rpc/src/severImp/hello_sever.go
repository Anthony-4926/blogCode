package severImp

import (
	"04_bidirectional_stream_rpc/src/pb/hello"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

type HelloSeverImp struct {
	hello.UnimplementedHelloServiceServer
}

func (h HelloSeverImp) Hello(server hello.HelloService_HelloServer) error {
	for {
		req, err := server.Recv()
		// err为EOF表示客户端数据发送完毕
		if err == io.EOF {
			for i := 0; i < 10; i++ {
				msg := fmt.Sprintf("pos1 hello %d", i)
				server.Send(&hello.HelloResponse{Msg: msg})
			}
			break
		}
		if err != nil {
			logrus.Warnln(err)
		} else {
			logrus.Infoln(req)
		}
	}
	// 再这还可以发送数据给客户端
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("pos2 hello %d", i)
		server.Send(&hello.HelloResponse{Msg: msg})
	}
	// 返回nil表示正常结束,服务端发完数据了
	return nil
}
