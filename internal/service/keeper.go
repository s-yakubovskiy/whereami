package service

import (
	"context"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/entity"
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

func (s *LocationKeeperService) StoreLocation(ctx context.Context, req *pb.StoreLocationRequest) (*pb.StoreLocationResponse, error) {
	err := s.uc.StoreLocation(ctx, ConvertToStructLocation(req.Location))
	return nil, err
}

func ConvertToStructLocation(protoLocation *pb.Location) *entity.Location {
	return &entity.Location{
		IP:          protoLocation.Ip,
		Country:     protoLocation.Country,
		CountryCode: protoLocation.CountryCode,
		Region:      protoLocation.Region,
		RegionCode:  protoLocation.RegionCode,
		City:        protoLocation.City,
		Timezone:    protoLocation.Timezone,
		Zip:         protoLocation.Zip,
		Flag:        protoLocation.Flag,
		Isp:         protoLocation.Isp,
		Asn:         protoLocation.Asn,
		Latitude:    protoLocation.Latitude,
		Longitude:   protoLocation.Longitude,
		Date:        protoLocation.Date,
		Vpn:         protoLocation.Vpn,
		Comment:     protoLocation.Comment,
		Scores: entity.LocationScores{
			FraudScore:  int(protoLocation.Scores.FraudScore),
			IsCrawler:   protoLocation.Scores.IsCrawler,
			Host:        protoLocation.Scores.Host,
			Proxy:       protoLocation.Scores.Proxy,
			VPN:         protoLocation.Scores.Vpn,
			Tor:         protoLocation.Scores.Tor,
			RecentAbuse: protoLocation.Scores.RecentAbuse,
			BotStatus:   protoLocation.Scores.BotStatus,
		},
		Gps: entity.GPSReport{
			Latitude:  protoLocation.Gps.Latitude,
			Longitude: protoLocation.Gps.Longitude,
			Altitude:  protoLocation.Gps.Altitude,
			Url:       protoLocation.Gps.Url,
		},
		Map: protoLocation.Map,
	}
}
