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
	options     []grpc.UnaryServerInterceptor
	skipMethods []string
}

func (i *interceptor) skipLog(method string) bool {
	//quick good lookup function
	for _, m := range i.skipMethods {
		if m == method {
			return true
		}
	}
	return false
}

type InterceptorOption func(in *interceptor)

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

func WithSkipMethod(methods []string) InterceptorOption {
	return func(in *interceptor) {
		in.skipMethods = methods
	}
}

func NewInterceptor(service string, logger log.Logger, options ...InterceptorOption) Interceptor {

	in := &interceptor{}

	//apply default interceptors
	WithInterecptor(kitgrpc.Interceptor)(in)
	WithInterecptor(loggingInterceptor(in, logger))(in)
	WithInterecptor(recoveryInterceptor(in, logger))(in)

	for _, option := range options {
		option(in)
	}
	return in
}
