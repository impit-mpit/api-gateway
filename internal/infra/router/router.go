package router

import (
	"context"
	"net/http"
	"neuro-most/api-gateway/config"
	"neuro-most/api-gateway/internal/adapters/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Router struct {
	cfg config.Config
}

func NewRouter(cfg config.Config) Router {
	return Router{
		cfg: cfg,
	}
}

type HTTPError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (r Router) Listen() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	r.SetupServices(ctx, mux, opts...)

	http.ListenAndServe(":8080", mux)
}

func (r Router) SetupServices(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	newsService := service.NewNewsService(r.cfg)
	newsService.RegisterNewsService(ctx, mux, opts...)
	mediaService := service.NewMediaService(r.cfg)
	mediaService.RegisterMediaService(ctx, mux, opts...)
	aiService := service.NewAIService(r.cfg)
	aiService.RegisterAIService(ctx, mux, opts...)
	tagService := service.NewTagsService(r.cfg)
	tagService.RegisterTagsService(ctx, mux, opts...)
	categoryService := service.NewCategoryService(r.cfg)
	categoryService.RegisterCategoryService(ctx, mux, opts...)
	authService := service.NewAuthService(r.cfg)
	authService.RegisterAuthService(ctx, mux, opts...)
	specService := service.NewSpecialistService(r.cfg)
	specService.RegisterSpecialistService(ctx, mux, opts...)
}
