package service

import (
	"context"
	"neuro-most/api-gateway/config"
	authv1 "neuro-most/api-gateway/gen/go/proto/auth/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type AuthService struct {
	cfg config.Config
}

func NewAuthService(cfg config.Config) AuthService {
	return AuthService{
		cfg: cfg,
	}
}

func (s AuthService) RegisterAuthService(ctx context.Context, mux *runtime.ServeMux, opts ...grpc.DialOption) {
	authv1.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, s.cfg.AuthService, opts)
}
