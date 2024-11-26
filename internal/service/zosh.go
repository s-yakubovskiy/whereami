package service

import (
	"context"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
)

// This line ensures at compile time that LocatorService implements the v1.WhrmiServiceServer interface.
var _ pb.ZoshServiceServer = &ZoshService{}

type ZoshService struct {
	pb.UnsafeZoshServiceServer
	// uc *locator.UseCase
}

// func NewZoshService(useCase *locator.UseCase) *LocationShowService {
func NewZoshService() *ZoshService {
	return &ZoshService{}
}

func (s *ZoshService) Live(ctx context.Context, req *pb.LiveRequest) (*pb.LiveResponse, error) {
	return &pb.LiveResponse{Message: "Up'n'Running, bro"}, nil
}
