package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/netbookai/log"
	"google.golang.org/grpc"
)

func loggingInterceptor(logger log.Logger) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		// log request and response data

		begin := time.Now()
		request := fmt.Sprintf("%+v", req)
		method := getMethod(info)

		logger.Debug(ctx, method, "method", method, "request", request)
		resp, err := handler(ctx, req)
		logger.Info(ctx, method, "method", method, "request", request, "response", resp, "error", err, "took", time.Since(begin))
		return resp, err
	}
}
