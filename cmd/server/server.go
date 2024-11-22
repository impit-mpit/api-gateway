package main

import (
	"neuro-most/api-gateway/config"
	"neuro-most/api-gateway/internal/infra/router"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	router := router.NewRouter(cfg)
	router.Listen()
}
