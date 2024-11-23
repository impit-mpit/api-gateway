package service

import (
	"context"
	"neuro-most/api-gateway/config"
	aiv1 "neuro-most/api-gateway/gen/go/proto/ai/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type AIService struct {
	cfg config.Config
}

func NewAIService(cfg config.Config) AIService {
	return AIService{
		cfg: cfg,
	}
}

func (s AIService) RegisterAIService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	aiv1.RegisterAIServiceHandlerFromEndpoint(ctx, mux, s.cfg.AIService, opts)
}
