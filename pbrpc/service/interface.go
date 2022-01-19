package service

const (
	Name = "HelloService"
)

type HelloService interface {
	Hello(req *Request, resp *Response) error
}