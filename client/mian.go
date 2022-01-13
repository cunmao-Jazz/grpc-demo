package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/cunmao-Jazz/grpc-demo/service"
)

var _ service.Service = (*HelloServiceClient)(nil)

//客户端构造函数
func NewHelloServiceClient(network, address string) (service.Service, error) {
	//首先是通过rpc.Dial拨号rpc服务，建立连接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	//客户端实现了基于JSON的编解码
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(name string, resp *string) error {
	// 然后通过client.Call调用具体的RPC方法
	// 在调用client.Call时:
	// 		第一个参数是用点号链接的RPC服务名字和方法名字，
	// 		第二个参数是 请求参数
	//      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值

	return c.client.Call(service.Name+".Hello", name, resp)
}

func main() {
	client, err := NewHelloServiceClient("tcp", "127.0.0.1:3580")
	if err != nil {
		panic(err)
	}
	var resp string
	client.Hello("boo", &resp)

	fmt.Println(resp)
}
