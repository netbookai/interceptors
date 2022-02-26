package interceptors

import (
	"strings"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func getMethod(info *grpc.UnaryServerInfo) string {
	splits := strings.Split(info.FullMethod, "/")
	return splits[len(splits)-1]
}

//GetInterceptors return the unary interceptos, with injected
// kit Interceptor - inject method name to the context
// loggingInterceptor - log request and response data, duration of the call
// recoveryInterceptor - recover from any API panics gracefully and logs error
func GetInterceptors(service string, logger *zap.SugaredLogger) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		kitgrpc.Interceptor,
		loggingInterceptor(service, logger),
		recoveryInterceptor(service, logger),
	)
}
