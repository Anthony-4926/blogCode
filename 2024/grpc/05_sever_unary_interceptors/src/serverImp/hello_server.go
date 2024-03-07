package serverImp

import (
	Hello "05_sever_unary_interceptors/src/pb/hello"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

type HelloServer struct {
	Hello.UnimplementedHelloServiceServer
}

func (h HelloServer) Hello(ctx context.Context, request *Hello.HelloRequest) (*Hello.HelloResponse, error) {
	if request == nil {
		logrus.Warnln("request is nil")
		return nil, errors.New("request is nil")
	}
	logrus.Info("request: ", request)

	resMsg := &Hello.HelloResponse{
		Who: &Hello.People{
			Name: "Tom",
		},
		Msg: "Hello " + request.GetWho().GetName(),
	}
	time.Sleep(time.Second)
	return resMsg, nil
}

// 以下是中间件
type UnaryMiddleFunc = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)

func TimeCounsumeMiddleware() UnaryMiddleFunc {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		start := time.Now()

		//-------------------上边是调用服务前的处理逻辑-------------------------------

		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		//-------------------下边是调用服务后的处理逻辑-----------------------------

		timeCount := time.Since(start)

		logrus.Info("TimeCount: ", timeCount)

		return res, nil
	}
}
