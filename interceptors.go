package interceptors

import (
	"strings"

	"github.com/netbookai/log"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func getMethod(info *grpc.UnaryServerInfo) string {
	splits := strings.Split(info.FullMethod, "/")
	return splits[len(splits)-1]
}

type Interceptor interface {
	Get() grpc.ServerOption
}

type interceptor struct {
	options []grpc.UnaryServerInterceptor
}

type InterceptorOption func(in *interceptor)

func NewInterceptor(service string, logger log.Logger, options ...InterceptorOption) Interceptor {

	in := &interceptor{}

	//apply default interceptors
	WithInterecptor(kitgrpc.Interceptor)(in)
	WithInterecptor(loggingInterceptor(logger))(in)
	WithInterecptor(recoveryInterceptor(logger))(in)

	for _, option := range options {
		option(in)
	}
	return in
}

func WithInterecptor(userInterceptor grpc.UnaryServerInterceptor) InterceptorOption {

	return func(in *interceptor) {
		in.options = append(in.options, userInterceptor)
	}

}

//Get return the unary interceptos composing of all default and user options
func (in *interceptor) Get() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		in.options...,
	)
}
