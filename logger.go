package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/netbookai/log"
	"google.golang.org/grpc"
)

func loggingInterceptor(in *interceptor, logger log.Logger) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		// log request and response data

		begin := time.Now()
		request := fmt.Sprintf("%+v", req)
		method := getMethod(info)

		skip := in.skipLog(method)
		msg := fmt.Sprintf("call to %s", method)

		args := []interface{}{msg, "method", method}
		if !skip {
			args = append(args, "request", request)
		}

		logger.Info(ctx, args...)
		resp, err := handler(ctx, req)

		args = []interface{}{msg, "method", method, "error", err, "took", time.Since(begin)}
		if !skip {
			args = append(args, "request", request, "response", resp)
		}

		logger.Info(ctx, args...)
		return resp, err
	}
}
