package router

import (
	"context"
	"net/http"
	"neuro-most/api-gateway/config"
	"neuro-most/api-gateway/internal/adapters/service"
	"neuro-most/api-gateway/internal/utils"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			md := make(map[string]string)
			if auth := req.Header.Get("Authorization"); auth != "" {
				md["authorization"] = auth
			}
			return metadata.New(md)
		}),
		// Дефолтный JSON маршалер для обычных запросов
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		// SSE маршалер только для streaming endpoints
		runtime.WithMarshalerOption("text/event-stream", utils.NewSSEMarshaler()),
		// Настройка заголовков только для streaming ответов
		runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, resp protoreflect.ProtoMessage) error {
			// Проверяем, является ли это streaming ответом
			if _, ok := resp.(interface{ GetMessage() string }); ok &&
				w.Header().Get("Content-Type") == "text/event-stream" {
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")
				w.Header().Set("Access-Control-Allow-Origin", "*")

				flusher, ok := w.(http.Flusher)
				if ok {
					flusher.Flush()
				}
			}
			return nil
		}),
	)
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
}
