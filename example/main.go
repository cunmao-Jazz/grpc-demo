package main

import (
	"fmt"

	"github.com/cunmao-Jazz/grpc-demo/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	s1 := &protobuf.String{
		Value: "wenqirui",
	}

	payload,err:=proto.Marshal(s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload)

	s2 := &protobuf.String{}

	proto.Unmarshal(payload,s2)

	fmt.Println(s2)

	sm := &protobuf.SampleMessage{
		TestOneof: &protobuf.SampleMessage_Sub1{
			Sub1: &protobuf.Sub1{Name: "test"},
		},
	}

	sm.GetSub1()
	sm.GetSub2()
	fmt.Println(sm.GetSub1())
	fmt.Println(sm.GetSub2())

	//any 使用

	//把sub1转化为any类型
	any,_:= anypb.New(&protobuf.Sub1{Name: "test"})
	errstat := &protobuf.ErrorStatus{
		Message: "Run",
		Details: any,
	}

	fmt.Println(errstat)

	//定义需要反序列化的空结构体
	foo := &protobuf.Sub1{}
	if err := anypb.UnmarshalTo(errstat.Details,foo,proto.UnmarshalOptions{});err != nil {
		fmt.Println(err)
	}
	fmt.Println(foo)

	fmt.Println("-----------------------------------------------------")

	//通过类型断言获取类型,errstat.Details.TypeUrl
	m,err := anypb.UnmarshalNew(errstat.Details,proto.UnmarshalOptions{})
	if err != nil {
		fmt.Println(err)
	}
	switch m.(type) {
	case *protobuf.Sub1:
		fmt.Println(m)
	default:
		fmt.Println("err")
	}
}