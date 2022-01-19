package main

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/cunmao-Jazz/grpc-demo/grpc/auth"
	"github.com/cunmao-Jazz/grpc-demo/grpc/protocol"
	"google.golang.org/grpc"
)

type Service struct {
	//HelloServiceServer 接口的默认实现,这个UnimplementedHelloServiceServer 默认帮你实现了Hello方法，如果你没有实现Hello方法,会使用他实现的默认Hello
	protocol.UnimplementedHelloServiceServer
}

//server端调用简单实现
func (s *Service) Hello(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	return &protocol.Response{
		Value: "hello:" + req.Value,
	}, nil
}

//服务端  ---stream ---->client
func (s *Service) Chat(stream protocol.HelloService_ChatServer) error {
	//可以把stream当成一个channel对象,io pipe
	for {
		// 用于接受client请求
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("char exit")
				return nil
			}
			return err
		}
		//用于处理服务端的响应
		if err := stream.Send(&protocol.Response{
			Value: "hello:" + req.Value,
		}); err != nil {
			fmt.Println(err)
		}
	}

}

func main() {
	// 如何把Service 作为一个 rpc暴露出去,提供服务
	server := grpc.NewServer(
		//请求响应模式的 认证中间件
		grpc.UnaryInterceptor(auth.GrpcAuthUnaryServerInterceptor()),
		//stream模式的认证中间件
		grpc.StreamInterceptor(auth.GrpcAuthStreamServerInterceptor()),		
	)

	//Service 申城的代码里面，提供了注册函数,把自己注册到grpc的server内
	protocol.RegisterHelloServiceServer(server, new(Service))

	// grpc.ChainUnaryInterceptor()

	//监听端口
	ls, err := net.Listen("tcp", "127.0.0.1:3580")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(ls); err != nil {
		fmt.Println(err)
	}
}
