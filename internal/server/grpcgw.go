package server

import (
	"context"
	"net"
	"net/http"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/service"
	"github.com/s-yakubovskiy/whereami/pkg/shudralogs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	metricsHandler = "/metrics"
)

// GrpcGatewaySrv wraps both the gRPC server and HTTP gateway
type GrpcGatewaySrv struct {
	grpcServer   *grpc.Server
	httpListener net.Listener
	grpcListener net.Listener
	logger       shudralogs.Logger
	httpAddress  string
	grpcAddress  string
}

// NewGrpcGatewaySrv initializes both the gRPC server and HTTP gateway
func NewGrpcGatewaySrv(
	cfg *config.Server,
	logger shudralogs.Logger,
	zoshService *service.ZoshService,
	locationService *service.LocationShowService,
) (*GrpcGatewaySrv, error) {
	// gRPC Listener
	grpcLis, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		logger.Errorf("failed to listen on gRPC address %s: %v", cfg.GRPC.Address, err)
		return nil, err
	}

	// HTTP Listener
	httpLis, err := net.Listen("tcp", cfg.HTTP.Address)
	if err != nil {
		logger.Errorf("failed to listen on HTTP address %s: %v", cfg.HTTP.Address, err)
		return nil, err
	}

	// Initialize Prometheus metrics for gRPC
	grpcMetrics := grpcprometheus.NewServerMetrics()

	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(
		grpcmiddleware.ChainUnaryServer(
			grpcMetrics.UnaryServerInterceptor(),  // Prometheus interceptor
			loggingUnaryServerInterceptor(logger), // Logging interceptor
		),
	))
	serverOptions = append(serverOptions, grpc.StreamInterceptor(
		grpcmiddleware.ChainStreamServer(
			grpcMetrics.StreamServerInterceptor(),
		),
	))

	grpcSrv := grpc.NewServer(serverOptions...)

	// Register gRPC services
	whrmi.RegisterZoshServiceServer(grpcSrv, zoshService)
	whrmi.RegisterLocationServiceServer(grpcSrv, locationService)

	// Enable reflection
	reflection.Register(grpcSrv)

	// Initialize Prometheus metrics
	grpcMetrics.InitializeMetrics(grpcSrv)

	return &GrpcGatewaySrv{
		grpcServer:   grpcSrv,
		httpListener: httpLis,
		grpcListener: grpcLis,
		logger:       logger,
		httpAddress:  cfg.HTTP.Address,
		grpcAddress:  cfg.GRPC.Address,
	}, nil
}

// ServeSync starts the gRPC server synchronously
func (gs *GrpcGatewaySrv) Serve(ctx context.Context) error {
	// Start gRPC server
	go func() {
		gs.logger.Infof("Starting gRPC server on %s", gs.grpcListener.Addr())
		if err := gs.grpcServer.Serve(gs.grpcListener); err != nil {
			gs.logger.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Start HTTP gateway
	go func() {
		mux := runtime.NewServeMux()

		dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		// Register gRPC Gateway endpoints
		if err := whrmi.RegisterZoshServiceHandlerFromEndpoint(ctx, mux, gs.grpcAddress, dialOptions); err != nil {
			gs.logger.Fatalf("failed to register ZoshService HTTP gateway: %v", err)
		}
		if err := whrmi.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, gs.grpcAddress, dialOptions); err != nil {
			gs.logger.Fatalf("failed to register LocationService HTTP gateway: %v", err)
		}

		// Add Prometheus metrics handler
		httpMux := http.NewServeMux()
		httpMux.Handle(metricsHandler, promhttp.Handler())
		httpMux.Handle("/", mux)

		gs.logger.Infof("Starting HTTP gateway on %s", gs.httpListener.Addr())
		if err := http.Serve(gs.httpListener, httpMux); err != nil {
			gs.logger.Fatalf("HTTP gateway failed: %v", err)
		}
	}()

	<-ctx.Done()
	gs.logger.Infof("Stopping servers...")

	// Graceful stop
	gs.grpcServer.GracefulStop()
	return nil
}
