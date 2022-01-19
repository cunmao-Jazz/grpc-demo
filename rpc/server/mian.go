package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"github.com/cunmao-Jazz/grpc-demo/rpc/service"
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

//tcp rpc

// func main() {
// 	//把要提供的服务注册给rpc框架
// 	err := rpc.RegisterName("HelloService", new(HelloService))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 监听socket
// 	listener, err := net.Listen("tcp", ":3580")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			panic(err)
// 		}

// 		//前面都是tcp的知识，到这个rpc就接管了
// 		//因此 你可以认为rpc 帮我们封装消息到函数调用的这个逻辑
// 		//提升了工作效率，逻辑比较简洁
// 		//实现服务端的json编解码
// 		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
// 	}
// }

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *NewRPCReadWriteCloser {
	return &NewRPCReadWriteCloser{w, r.Body}
}

type NewRPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}
	// RPC的服务架设在“/jsonrpc”路径，
	// 在处理函数中基于http.ResponseWriter和http.Request类型的参数构造一个io.ReadWriteCloser类型的conn通道。
	// 然后基于conn构建针对服务端的json编码解码器。
	// 最后通过rpc.ServeRequest函数为每次请求处理一次RPC方法调用
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		//如何把HTTP的请求 --->
		//基于HTTP request body 和 Respones构造了一个Conn
		conn := NewRPCReadWriteCloserFromHTTP(w, r)
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":3580", nil)
}
