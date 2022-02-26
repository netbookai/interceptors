package interceptors

import (
	"testing"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

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
	serverOptions := GetInterceptors("my-service", zap.L().Sugar())

	if got, ok := serverOptions.(grpc.ServerOption); !ok {
		t.Fatalf("GetInterceptors :  type missmatch, expected 'grpc.ServerOption', got '%s'", got)
	}
}
