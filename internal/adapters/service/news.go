package service

import (
	"context"
	"neuro-most/api-gateway/config"
	newsv1 "neuro-most/api-gateway/gen/go/proto/news/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type NewsService struct {
	cfg config.Config
}

func NewNewsService(cfg config.Config) NewsService {
	return NewsService{
		cfg: cfg,
	}
}

func (s NewsService) RegisterNewsService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	newsv1.RegisterNewsServiceHandlerFromEndpoint(ctx, mux, s.cfg.NewsService, opts)
}
