package main

import (
	"fmt"
	"net"
	"net/rpc"
	"github.com/cunmao-Jazz/grpc-demo/service"
)

//var
//var var1 int = "str"
// var _ Service = &HelloService{}
//我们声明了一个空指针,强制把这个指针转换成了一个HelloService
var _ service.Service = (*HelloService)(nil)

type HelloService struct {
}

//业务场景
//该函数需要被客户端调用
//改造成符合rpc规范的函数签名
//1 第一个参数 requset，inerface{}
func (s *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("Hello,%s", req)
	return nil
}

func main() {
	//把要提供的服务注册给rpc框架
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}

	// 监听socket
	listener, err := net.Listen("tcp", ":3580")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		//前面都是tcp的知识，到这个rpc就接管了
		//因此 你可以认为rpc 帮我们封装消息到函数调用的这个逻辑
		//提升了工作效率，逻辑比较简洁
		go rpc.ServeConn(conn)
	}
}
