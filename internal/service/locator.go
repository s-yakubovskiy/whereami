package service

import (
	"context"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
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
	err := s.uc.ShowLocation(ctx)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (s *LocationShowService) Init(ctx context.Context, req *pb.InitRequest) (*pb.InitResponse, error) {
	return nil, nil
}
