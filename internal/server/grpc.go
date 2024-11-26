package server

import (
	"context"
	"net"
	"net/http"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/logging"
	"github.com/s-yakubovskiy/whereami/internal/metrics"
	"github.com/s-yakubovskiy/whereami/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GrpcSrv wraps the gRPC server and its configuration
type GrpcSrv struct {
	grpcServer *grpc.Server
	listener   net.Listener
	logger     logging.Logger
}

// NewGrpcSrv initializes the gRPC server
func NewGrpcSrv(
	cfg *config.Server,
	logger logging.Logger,
	zoshService *service.ZoshService, // Replace with actual ZoshService implementation
	locationService *service.LocationShowService, // Replace with actual LocationService implementation
) (*GrpcSrv, error) {
	lis, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		logger.Errorf("failed to listen on address %s: %v", cfg.GRPC.Address, err)
		return nil, err
	}

	// Initialize Prometheus gRPC metrics
	grpcMetrics := grpcprometheus.NewServerMetrics()
	metrics.RegisterMetrics() // Register custom metrics

	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(
		grpcmiddleware.ChainUnaryServer(
			grpcMetrics.UnaryServerInterceptor(),  // Prometheus interceptor
			loggingUnaryServerInterceptor(logger), // Logging interceptor
			customMetricsInterceptor(logger),      // Custom metrics interceptor
		),
	))

	serverOptions = append(serverOptions, grpc.StreamInterceptor(
		grpcmiddleware.ChainStreamServer(
			grpcMetrics.StreamServerInterceptor(), // Prometheus interceptor for streams
		),
	))

	srv := grpc.NewServer(serverOptions...)
	reflection.Register(srv)

	// Register your gRPC services here
	whrmi.RegisterZoshServiceServer(srv, zoshService)
	whrmi.RegisterLocationServiceServer(srv, locationService)

	// Initialize default Prometheus gRPC metrics
	grpcMetrics.InitializeMetrics(srv)

	// Start Prometheus HTTP server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logger.Infof("Metrics server running on %s", cfg.Metrics.Address)
		if err := http.ListenAndServe(cfg.Metrics.Address, nil); err != nil {
			logger.Fatalf("Metrics server failed: %v", err)
		}
	}()

	return &GrpcSrv{
		grpcServer: srv,
		listener:   lis,
		logger:     logger,
	}, nil
}

// ServeSync starts the gRPC server synchronously
func (cs *GrpcSrv) ServeSync() error {
	cs.logger.Infof("Starting gRPC server on %s", cs.listener.Addr())
	if err := cs.grpcServer.Serve(cs.listener); err != nil {
		cs.logger.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

// Serve starts the gRPC server with graceful shutdown
func (cs *GrpcSrv) Serve(ctx context.Context) error {
	go func() {
		cs.logger.Infof("Starting gRPC server on %s", cs.listener.Addr())
		if err := cs.grpcServer.Serve(cs.listener); err != nil {
			cs.logger.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	cs.logger.Infof("Stopping gRPC server...")
	cs.grpcServer.GracefulStop()
	return nil
}

// Sample UnaryServerInterceptor implementation for logging
func loggingUnaryServerInterceptor(logger logging.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Infof("Received request for method: %s", info.FullMethod)
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Errorf("Method %s encountered error: %v", info.FullMethod, err)
		}
		// Logging or other middleware logic comes here...
		return resp, err
	}
}
