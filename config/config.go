package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	NewsService     string `env:"NEWS_SERVICE" env-default:"localhost:50051"`
	MediaService    string `env:"MEDIA_SERVICE" env-default:"localhost:50052"`
	AIService       string `env:"AI_SERVICE" env-default:"localhost:50053"`
	TagsService     string `env:"TAGS_SERVICE" env-default:"localhost:50054"`
	CategoryService string `env:"CATEGORY_SERVICE" env-default:"localhost:50055"`
	AuthService     string `env:"AUTH_SERVICE" env-default:"localhost:50056"`
}

func NewLoadConfig() (Config, error) {
	var cfg Config
	cleanenv.ReadConfig(".env", &cfg)
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
