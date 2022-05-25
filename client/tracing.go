package interceptors

import (
	"context"

	"github.com/google/uuid"
	"github.com/netbookai/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const TRACE_ID = "trace-id"

func TracingInterceptor(logger log.Logger) grpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		traceid, ok := ctx.Value(TRACE_ID).(string)
		if !ok {
			logger.Warn(ctx, "invalid traceid found in context", "traceid", ctx.Value(TRACE_ID))
			traceid = generateTraceId()
			logger.Info(ctx, "generated new trace id", "traceid", traceid)
		}

		ctx = metadata.AppendToOutgoingContext(ctx, TRACE_ID, traceid)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func generateTraceId() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
