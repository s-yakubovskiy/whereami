package service

import (
	"context"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/entity"
	"github.com/s-yakubovskiy/whereami/internal/usecase/locator"
)

// This line ensures at compile time that LocatorService implements the v1.WhrmiServiceServer interface.
var _ pb.LocationServiceServer = &LocationShowService{}

type LocationShowService struct {
	pb.UnsafeLocationServiceServer
	uc *locator.UseCase
}

func NewLocationShowService(useCase *locator.UseCase) *LocationShowService {
	return &LocationShowService{uc: useCase}
}

func (s *LocationShowService) ShowLocation(ctx context.Context, req *pb.ShowLocationRequest) (*pb.ShowLocationResponse, error) {
	data, err := s.uc.ShowLocation(ctx)
	if err != nil {
		return nil, err
	}
	// return nil, err
	return &pb.ShowLocationResponse{
		Location: ConvertToProtoLocation(data),
	}, nil
}

func (s *LocationShowService) Init(ctx context.Context, req *pb.InitRequest) (*pb.InitResponse, error) {
	return nil, nil
}

func (s *LocationShowService) GetLocation(ctx context.Context, req *pb.GetLocationRequest) (*pb.GetLocationResponse, error) {
	data, err := s.uc.GetLocation(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetLocationResponse{
		Location: ConvertToProtoLocation(data),
	}, nil
}

func ConvertToProtoLocation(location *entity.Location) *pb.Location {
	return &pb.Location{
		Ip:          location.IP,
		Country:     location.Country,
		CountryCode: location.CountryCode,
		Region:      location.Region,
		RegionCode:  location.RegionCode,
		City:        location.City,
		Timezone:    location.Timezone,
		Zip:         location.Zip,
		Flag:        location.Flag,
		Isp:         location.Isp,
		Asn:         location.Asn,
		Latitude:    location.Latitude,
		Longitude:   location.Longitude,
		Date:        location.Date,
		Vpn:         location.Vpn,
		Comment:     location.Comment,
		Scores: &pb.LocationScores{
			FraudScore:  int32(location.Scores.FraudScore),
			IsCrawler:   location.Scores.IsCrawler,
			Host:        location.Scores.Host,
			Proxy:       location.Scores.Proxy,
			Vpn:         location.Scores.VPN,
			Tor:         location.Scores.Tor,
			RecentAbuse: location.Scores.RecentAbuse,
			BotStatus:   location.Scores.BotStatus,
		},
		Gps: &pb.GPSReport{
			Latitude:  location.Gps.Latitude,
			Longitude: location.Gps.Longitude,
			Altitude:  location.Gps.Altitude,
			Url:       location.Gps.Url,
		},
		Map: location.Map,
	}
}
