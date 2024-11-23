package service

import (
	"context"
	"neuro-most/api-gateway/config"
	mediav1 "neuro-most/api-gateway/gen/go/proto/media/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type MediaService struct {
	cfg config.Config
}

func NewMediaService(cfg config.Config) MediaService {
	return MediaService{
		cfg: cfg,
	}
}

func (s MediaService) RegisterMediaService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	mediav1.RegisterMediaServiceHandlerFromEndpoint(ctx, mux, s.cfg.MediaService, opts)
}
