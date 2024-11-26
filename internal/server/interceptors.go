package server

import (
	"context"

	"github.com/s-yakubovskiy/whereami/internal/logging"
	"google.golang.org/grpc"
)

// Sample UnaryServerInterceptor implementation for logging
func loggingUnaryServerInterceptor(logger logging.Logger) grpc.UnaryServerInterceptor {
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
