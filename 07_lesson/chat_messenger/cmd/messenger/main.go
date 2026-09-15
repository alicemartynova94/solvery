package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"net"
	"net/http"
	"os"
	"os/signal"
	grpcsrv "solvery/07_lesson/chat_messenger/internal/handler/grpc"
	httpsrv "solvery/07_lesson/chat_messenger/internal/handler/http"
	"solvery/07_lesson/chat_messenger/internal/logger"
	"solvery/07_lesson/chat_messenger/pkg/messenger/v1"
	"solvery/07_lesson/chat_messenger/pkg/openapi"
	"strconv"
	"syscall"
	"time"
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

func main() {
	cfg := MustLoadConfig()

	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = log.Sync()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	errCh := make(chan error, 2)

	db, err := ConnectDB(ctx, cfg)
	if err != nil {
		log.Error("db connection failed", zap.Error(err))
		return
	}
	defer db.Close()
	httpServer := NewHTTPServer(cfg)
	grpcServer := NewGRPCServer()

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("http server: %w", err)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", net.JoinHostPort(cfg.GRPC.Host, strconv.Itoa(cfg.GRPC.Port)))
		if err != nil {
			errCh <- fmt.Errorf("grpc listen: %w", err)
			return
		}
		defer listener.Close()

		err = grpcServer.Serve(listener)
		if err != nil {
			errCh <- fmt.Errorf("grpc server: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		log.Error("server failed", zap.Error(err))
	case <-ctx.Done():
		log.Info("shutdown signal received")
	}

	gracefulShutdown(log, httpServer, grpcServer)
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

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}

func NewHTTPServer(cfg *Config) *http.Server {
	mux := http.NewServeMux()
	apiServer := httpsrv.NewHandler()
	openapi.HandlerFromMux(apiServer, mux)

	swaggerFS := http.FileServer(http.Dir(cfg.Swagger.Path))
	mux.Handle("GET /swagger/", http.StripPrefix("/swagger/", swaggerFS))
	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    net.JoinHostPort(cfg.HTTP.Host, strconv.Itoa(cfg.HTTP.Port)),
		Handler: mux,
	}
}

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	messenger.RegisterMessengerServiceServer(server, grpcsrv.NewHandler())

	return server
}

func gracefulShutdown(log *zap.Logger, httpServer *http.Server, grpcServer *grpc.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Error("http shutdown failed", zap.Error(err))
	}

	grpcDone := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(grpcDone)
	}()

	select {
	case <-grpcDone:
		log.Info("gRPC server stopped")

	case <-ctx.Done():
		log.Warn("gRPC graceful shutdown timeout", zap.Error(err))
		grpcServer.Stop()
	}
}

func ConnectDB(ctx context.Context, cfg *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
		cfg.Database.ConnectTimeout,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("pg connection pull: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.Database.ConnectTimeout)*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("db ping: %w", err)
	}

	return db, nil
}
