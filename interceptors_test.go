package interceptors

import (
	"context"
	"testing"

	"github.com/netbookai/log"
	"github.com/netbookai/log/loggers/zap"
	"google.golang.org/grpc"
)

func customInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {

		return nil, nil
	}
}

func Test_getMethod(t *testing.T) {
	info := grpc.UnaryServerInfo{
		FullMethod: "service/GetCost",
	}

	expected := "GetCost"
	got := getMethod(&info)
	if got != expected {
		t.Fatalf("getMethod : expected %s, got %s", expected, got)
	}
}

func Test_GetInterceptors(t *testing.T) {
	//we will assert the types for now
	//as there is no build, we can find any compilation issue here

	//with default options
	logger := log.NewLogger(zap.NewLogger())
	ints := NewInterceptor("my-service", logger)

	serverOptions := ints.Get()

	if got, ok := serverOptions.(grpc.ServerOption); !ok {

		t.Fatalf("GetInterceptors :  type missmatch, expected 'grpc.ServerOption', got '%s'", got)
	}

	//with default options
	ints = NewInterceptor("my-service", logger, WithInterecptor(customInterceptor()))

	serverOptions = ints.Get()

	if got, ok := serverOptions.(grpc.ServerOption); !ok {

		t.Fatalf("GetInterceptors :  type missmatch, expected 'grpc.ServerOption', got '%s'", got)
	}
}

func Test_skipLog(t *testing.T) {
	in := &interceptor{}
	WithSkipMethod([]string{"password"})(in)

	if !in.skipLog("password") {
		t.Fatalf("skipLog: expected : true, Got : false")
	}

	if in.skipLog("ping") {
		t.Fatalf("skipLog: expected : false, Got :true")
	}
}
