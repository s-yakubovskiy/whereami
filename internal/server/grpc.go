package server

import (
	"context"
	"net"

	"github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/logging"
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

// NewGrpcSrv initializes the gRPC server and returns a CustomServer instance
func NewGrpcSrv(
	cfg *config.Server,
	logger logging.Logger,
	zoshService *service.ZoshService,
) (*GrpcSrv, error) {
	lis, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		logger.Errorf("failed to listen on address %s: %w", cfg.GRPC.Address, err)
		return nil, err
	}
	logger.Debugf("gRPC server initially setup on %s", cfg.GRPC.Address)

	var serverOptions []grpc.ServerOption
	unaryInterceptor := loggingUnaryServerInterceptor(logger)
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(unaryInterceptor))

	srv := grpc.NewServer(serverOptions...)
	reflection.Register(srv)

	whrmi.RegisterZoshServiceServer(srv, zoshService)

	return &GrpcSrv{
		grpcServer: srv,
		listener:   lis,
		logger:     logger,
	}, nil
}

// Serve starts the gRPC server; should be called after server is initialized
func (cs *GrpcSrv) ServeSync() error {
	cs.logger.Printf("Starting gRPC server serve process...")
	if err := cs.grpcServer.Serve(cs.listener); err != nil {
		cs.logger.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

func (cs *GrpcSrv) Serve(ctx context.Context) error {
	go func() {
		cs.logger.Printf("Starting gRPC server serve process...")

		if err := cs.grpcServer.Serve(cs.listener); err != nil {
			cs.logger.Fatalf("failed to serve: %v", err)
		}
	}()

	// Listen for cancellation and shut down gracefully
	<-ctx.Done()
	cs.logger.Printf("Stopping gRPC server...")
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
