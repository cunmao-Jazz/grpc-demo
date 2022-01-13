package service

const (
	Name = "HelloService"
)

type Service interface{
	Hello(req string,resp *string)error
}

