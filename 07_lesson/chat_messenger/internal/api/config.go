package api

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App struct {
		Env string `yaml:"env"`
	} `yaml:"app"`

	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"http"`

	Swagger struct {
		Path string `yaml:"path"`
	} `yaml:"swagger"`

	GRPC struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"grpc"`

	Database struct {
		Host           string `yaml:"host"`
		Port           int    `yaml:"port"`
		User           string `yaml:"user"`
		Password       string `yaml:"password"`
		Name           string `yaml:"name"`
		SSLMode        string `yaml:"sslmode"`
		ConnectTimeout int    `yaml:"connect_timeout"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	path := fmt.Sprintf("configs/%s.yaml", env)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file: %w", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("yaml decoder: %w", err)
	}

	return &cfg, nil
}
