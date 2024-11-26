package server

import (
	"context"
	"time"

	"github.com/s-yakubovskiy/whereami/internal/logging"
	"github.com/s-yakubovskiy/whereami/internal/metrics"
	"google.golang.org/grpc"
)

// customMetricsInterceptor tracks custom metrics for gRPC methods
func customMetricsInterceptor(logger logging.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		// Record custom metrics
		status := "ok"
		if err != nil {
			status = "error"
		}
		metrics.RequestCounter.WithLabelValues(info.FullMethod, status).Inc()
		metrics.RequestLatency.WithLabelValues(info.FullMethod).Observe(time.Since(start).Seconds())

		if err != nil {
			logger.Errorf("Method %s encountered error: %v", info.FullMethod, err)
		}
		return resp, err
	}
}
