package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	HttpTimeout time.Duration `yaml:"http_timeout"`
	HttpPort    int           `yaml:"http_port"`
}

func ReadConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("config/dev.yaml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
