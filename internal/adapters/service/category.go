package service

import (
	"context"
	"neuro-most/api-gateway/config"
	categoryv1 "neuro-most/api-gateway/gen/go/proto/category/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type CategoryService struct {
	cfg config.Config
}

func NewCategoryService(cfg config.Config) CategoryService {
	return CategoryService{
		cfg: cfg,
	}
}

func (s CategoryService) RegisterCategoryService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	categoryv1.RegisterCategoryServiceHandlerFromEndpoint(ctx, mux, s.cfg.CategoryService, opts)
}
