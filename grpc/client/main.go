package main

import (
	"context"
	"fmt"
	"time"
	"github.com/cunmao-Jazz/grpc-demo/grpc/auth"
	"github.com/cunmao-Jazz/grpc-demo/grpc/protocol"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3580",
	 grpc.WithInsecure(),
	 grpc.WithPerRPCCredentials(auth.NewAuthentication("admin","123456")),
	)
	if err != nil {
		fmt.Println(err)
	}

	// //请求时,携带meta信息
	// md := metadata.New(map[string]string{auth.ClientHeaderKey:"admin",auth.ClientSecretKey:"12456"})

	// //客户端发送时 需要携带的上下文信息
	// //服务端接受时,对于服务端来说 这个就是InconmingContext   FromIncomingContext
	// ctx:=metadata.NewOutgoingContext(context.Background(),md)

	client := protocol.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &protocol.Request{Value: "wen"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	//返回一个channle io pipe
	stream,err:=client.Chat(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := stream.Send(&protocol.Request{
				Value: "waa",
			})
			if err != nil {
				fmt.Println(err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}()

	for {
		resp,err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	}
}
