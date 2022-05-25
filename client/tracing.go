package interceptors

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/netbookai/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const TRACE_ID = "trace-id"

func tracingInterceptor(logger log.Logger) grpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		traceid, ok := ctx.Value(TRACE_ID).(string)
		if !ok {
			logger.Warn(ctx, "invalid traceid found in context. generating fallback id", "traceid", ctx.Value(TRACE_ID))
			traceid = generateFallBackTraceId()
		}

		ctx = metadata.AppendToOutgoingContext(ctx, TRACE_ID, traceid)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func generateFallBackTraceId() string {
	id, _ := uuid.NewRandom()
	return fmt.Sprintf("fallback-%s", id.String())
}

func WithTracingInterceptor(logger log.Logger) grpc.DialOption {
	return grpc.WithUnaryInterceptor(tracingInterceptor(logger))
}
