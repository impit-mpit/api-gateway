package service

import (
	"context"
	"neuro-most/api-gateway/config"
	tagsv1 "neuro-most/api-gateway/gen/go/proto/tags"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type TagsService struct {
	cfg config.Config
}

func NewTagsService(cfg config.Config) TagsService {
	return TagsService{
		cfg: cfg,
	}
}

func (s TagsService) RegisterTagsService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	tagsv1.RegisterTagsServiceHandlerFromEndpoint(ctx, mux, s.cfg.TagsService, opts)
}
