package service

import (
	"context"
	"neuro-most/api-gateway/config"
	specialistv1 "neuro-most/api-gateway/gen/go/proto/specialist/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type SpecialistService struct {
	cfg config.Config
}

func NewSpecialistService(cfg config.Config) SpecialistService {
	return SpecialistService{
		cfg: cfg,
	}
}

func (s SpecialistService) RegisterSpecialistService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	specialistv1.RegisterSpecialistServiceHandlerFromEndpoint(ctx, mux, s.cfg.SpecialistService, opts)
}
