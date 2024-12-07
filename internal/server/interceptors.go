package server

import (
	"context"
	"reflect"

	"github.com/s-yakubovskiy/whereami/pkg/shudralogs"
	"google.golang.org/grpc"
)

// Sample UnaryServerInterceptor implementation for logging
func loggingUnaryServerInterceptor(logger shudralogs.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Debugf("Received request for method: %s", info.FullMethod)
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Errorf("Method %s encountered error: %v", info.FullMethod, err)
		}
		// Logging or other middleware logic comes here...
		return resp, err
	}
}

func structToMap(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()
	result := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		result[field.Name] = value
	}
	return result
}
