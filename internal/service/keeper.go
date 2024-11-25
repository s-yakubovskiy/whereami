package service

import (
	"context"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/usecase/keeper"
)

// This line ensures at compile time that LocatorService implements the v1.WhrmiServiceServer interface.
var _ pb.LocationKeeperServiceServer = &LocationKeeperService{}

type LocationKeeperService struct {
	pb.UnsafeLocationKeeperServiceServer
	uc *keeper.UseCase
}

func NewLocationKeeperService(useCase *keeper.UseCase) *LocationKeeperService {
	return &LocationKeeperService{uc: useCase}
}

func (s *LocationKeeperService) Init(ctx context.Context, req *pb.InitRequest) (*pb.InitResponse, error) {
	s.uc.InitDb(ctx)
	return nil, nil
}

func (s *LocationKeeperService) AddVpnInterface(ctx context.Context, req *pb.AddVpnInterfaceRequest) (*pb.AddVpnInterfaceResponse, error) {
	s.uc.AddVPNInterface(ctx, req.Vpninterface)
	return &pb.AddVpnInterfaceResponse{}, nil
}

func (s *LocationKeeperService) ListVpnInterfaces(ctx context.Context, req *pb.ListVpnInterfacesRequest) (*pb.ListVpnInterfacesResponse, error) {
	ifaces, err := s.uc.ListVPNInterfaces()
	return convertToListVpnInterfacesResponse(ifaces), err
}

func convertToListVpnInterfacesResponse(p []string) *pb.ListVpnInterfacesResponse {
	return &pb.ListVpnInterfacesResponse{
		Vpninterfaces: p,
	}
}

func (s *LocationKeeperService) ImportLocations(ctx context.Context, req *pb.ImportLocationsRequest) (*pb.ImportLocationsResponse, error) {
	return nil, s.uc.ImportLocations(ctx, req.Importpath)
}

func (s *LocationKeeperService) ExportLocations(ctx context.Context, req *pb.ExportLocationsRequest) (*pb.ExportLocationsResponse, error) {
	return nil, s.uc.ExportLocations(ctx, req.Exportpath)
}
