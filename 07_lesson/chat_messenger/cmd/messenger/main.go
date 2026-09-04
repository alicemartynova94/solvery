package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"net/http"
	"os"
	"solvery/07_lesson/chat_messenger/internal/api"
	"solvery/07_lesson/chat_messenger/internal/api/generated"
	"strconv"
)

type Config struct {
	App struct {
		Env string `yaml:"env"`
	} `yaml:"app"`

	HTTP struct {
		Enabled bool   `yaml:"enabled"`
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
	} `yaml:"http"`

	Swagger struct {
		Enabled bool   `yaml:"enabled"`
		Path    string `yaml:"path"`
	} `yaml:"swagger"`

	RPC struct {
		Enabled bool   `yaml:"enabled"`
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
	} `yaml:"grpc"`

	Database struct {
		Port           int    `yaml:"port"`
		User           string `yaml:"user"`
		Password       string `yaml:"password"`
		Name           string `yaml:"name"`
		SSLMode        string `yaml:"sslmode"`
		ConnectTimeout int    `yaml:"connect_timeout"`
	} `yaml:"database"`
}

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	cfg, err := LoadConfig(fmt.Sprintf("07_lesson/chat_messenger/configs/%s.yaml", env))
	if err != nil {
		log.Fatal(err)
	}

	httpServer := NewHTTPServer(cfg)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func NewHTTPServer(cfg *Config) *http.Server {
	mux := http.NewServeMux()

	apiServer := api.NewServer()
	generated.HandlerFromMux(apiServer, mux)

	// Swagger UI
	swaggerFS := http.FileServer(http.Dir("07_lesson/chat_messenger/api"))
	mux.Handle("GET /swagger/", http.StripPrefix("/swagger/", swaggerFS))

	return &http.Server{
		Addr:    net.JoinHostPort(cfg.HTTP.Host, strconv.Itoa(cfg.HTTP.Port)),
		Handler: mux,
	}
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
