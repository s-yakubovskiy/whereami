package service

import (
	"context"
	"encoding/json"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"

	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/usecase/zosher"
	"google.golang.org/protobuf/types/known/structpb"
)

// This line ensures at compile time that LocatorService implements the v1.WhrmiServiceServer interface.
var _ pb.ZoshServiceServer = &ZoshService{}

type ZoshService struct {
	pb.UnsafeZoshServiceServer
	uc *zosher.UseCase
}

// func NewZoshService(useCase *locator.UseCase) *LocationShowService {
func NewZoshService(uc *zosher.UseCase) *ZoshService {
	return &ZoshService{uc: uc}
}

func (s *ZoshService) Live(ctx context.Context, req *pb.LiveRequest) (*pb.LiveResponse, error) {
	return &pb.LiveResponse{Message: "Up'n'Running, bro"}, nil
}

func (s *ZoshService) Version(ctx context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	ver := s.uc.Version()
	return &pb.VersionResponse{
		Version: ver.Version,
		Commit:  ver.Commit,
		Date:    ver.Date,
	}, nil
}

func (s *ZoshService) Config(ctx context.Context, req *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	// Serialize the configuration to JSON
	configJSON, err := json.Marshal(config.Cfg)
	if err != nil {
		return nil, err
	}

	// Convert JSON to protobuf Struct
	var configMap map[string]interface{}
	if err := json.Unmarshal(configJSON, &configMap); err != nil {
		return nil, err
	}

	protoStruct, err := structpb.NewStruct(configMap)
	if err != nil {
		return nil, err
	}

	return &pb.ConfigResponse{
		Config: protoStruct,
	}, nil
}
