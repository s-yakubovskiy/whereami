package server

import (
	"context"
	"net"
	"net/http"
	"reflect"

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

	pb "github.com/s-yakubovskiy/zoshlib/api/zosh/v1"
	zl "github.com/s-yakubovskiy/zoshlib/service"
)

const (
	metricsHandler = "/metrics"
)

// GrpcSrv wraps the gRPC server and its configuration
type GrpcSrv struct {
	grpcServer *grpc.Server
	listener   net.Listener
	logger     logging.Logger
	address    string
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
			// customMetricsInterceptor(logger),      // Custom metrics interceptor
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

	zoshUseCase := zl.NewZoshUseCase(structToMap(&config.Cfg))
	zoshlibService := zl.NewZoshService(zoshUseCase)
	pb.RegisterZoshServiceServer(srv, zoshlibService)

	// Initialize default Prometheus gRPC metrics
	grpcMetrics.InitializeMetrics(srv)

	return &GrpcSrv{
		grpcServer: srv,
		listener:   lis,
		logger:     logger,
		address:    cfg.Metrics.Address,
	}, nil
}

// ServeSync starts the gRPC server synchronously
func (cs *GrpcSrv) ServeSync() error {
	go func() {
		http.Handle(metricsHandler, promhttp.Handler())
		cs.logger.Infof("Metrics server running on %s", cs.address)
		if err := http.ListenAndServe(cs.address, nil); err != nil {
			cs.logger.Fatalf("Metrics server failed: %v", err)
		}
	}()
	cs.logger.Infof("Starting gRPC server on %s", cs.listener.Addr())
	if err := cs.grpcServer.Serve(cs.listener); err != nil {
		cs.logger.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

// Serve starts the gRPC server with graceful shutdown
func (cs *GrpcSrv) Serve(ctx context.Context) error {
	// Start Prometheus HTTP server
	go func() {
		cs.logger.Infof("Starting gRPC server on %s", cs.listener.Addr())
		if err := cs.grpcServer.Serve(cs.listener); err != nil {
			cs.logger.Fatalf("failed to serve: %v", err)
		}
	}()
	go func() {
		http.Handle(metricsHandler, promhttp.Handler())
		cs.logger.Infof("Metrics server running on %s", cs.address)
		if err := http.ListenAndServe(cs.address, nil); err != nil {
			cs.logger.Fatalf("Metrics server failed: %v", err)
		}
	}()

	<-ctx.Done()
	cs.logger.Infof("Stopping gRPC server...")
	cs.grpcServer.GracefulStop()
	return nil
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
