package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

//通过暴露Auth函数,来提供grpc 中间件函数
func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return newGrpcAuther().Auth
}
//通过暴露StreamAuth函数,来grpc stream中间件函数
func GrpcAuthStreamServerInterceptor() grpc.StreamServerInterceptor {
	return newGrpcAuther().StreamAuth
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

type grpcAuther struct {
}

// 中间件
func (a *grpcAuther) Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	//从ctx 获取到了metadata信息
	md, _ := metadata.FromIncomingContext(ctx)

	fmt.Println(md)

	//认证请求的合法性
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)
	err = a.validateServiceCredential(clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	//请求认证通过,请求交给后面的handler调用后编写

	return handler(ctx, req)
}

//从metadata中获取客户端凭证
func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (clientId, clientSecret string) {
	cids := md.Get(ClientHeaderKey)
	sids := md.Get(ClientSecretKey)

	if len(cids) > 0 {
		clientId = cids[0]
	}

	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client credential %s,%s not provided", clientId, clientSecret)
	}

	if !(clientId == "admin" && clientSecret == "123456") {
		return status.Errorf(codes.Unauthenticated, "client-id or client-secret is not conrect")
	}

	return nil
}

func (a *grpcAuther) StreamAuth(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	//从Sream请求中获取上下文
	md, _ := metadata.FromIncomingContext(ss.Context())

	fmt.Println(md)

	//认证请求的合法性
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)
	err := a.validateServiceCredential(clientId, clientSecret)
	if err != nil {
		return err
	}

	//请求认证通过,请求交给后面的handler调用后编写
	return handler(srv, ss)

}
