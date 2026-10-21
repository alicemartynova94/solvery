package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	grpcsrv "solvery/07_lesson/chat_messenger/internal/handler/grpc"
	httpsrv "solvery/07_lesson/chat_messenger/internal/handler/http"
	"solvery/07_lesson/chat_messenger/internal/logger"
	"solvery/07_lesson/chat_messenger/internal/repository/memory"
	"solvery/07_lesson/chat_messenger/internal/repository/synchronization"
	"solvery/07_lesson/chat_messenger/internal/service"
	"solvery/07_lesson/chat_messenger/pkg/messenger/v1"
	"solvery/07_lesson/chat_messenger/pkg/openapi"
	"strconv"
	"syscall"
	"time"
)

func Run(cfg *Config) {
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

	mu := synchronization.NewMutex()
	sessionRepo := memory.NewSessionRepository(mu)
	chatRepo := memory.NewChatRepository()
	sessionSrv := service.NewSessionService(sessionRepo, cfg.Limits.MaxSessions)
	chatSrv := service.NewChatService(mu, chatRepo, sessionRepo, cfg.Limits.MaxChats)
	httpHandler := httpsrv.NewHandler(sessionSrv, chatSrv, log)
	grpcHandler := grpcsrv.NewHandler(sessionSrv, log)

	httpServer := NewHTTPServer(cfg, httpHandler)
	grpcServer := NewGRPCServer(grpcHandler)

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

func NewHTTPServer(cfg *Config, apiServer *httpsrv.Handler) *http.Server {
	mux := http.NewServeMux()
	openapi.HandlerFromMux(apiServer, mux)

	swaggerFS := http.FileServer(http.Dir(cfg.Swagger.Path))
	mux.Handle("GET /swagger/", http.StripPrefix("/swagger/", swaggerFS))
	openapiFS := http.FileServer(http.Dir("api/openapi"))
	mux.Handle("GET /openapi/", http.StripPrefix("/openapi/", openapiFS))
	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    net.JoinHostPort(cfg.HTTP.Host, strconv.Itoa(cfg.HTTP.Port)),
		Handler: mux,
	}
}

func NewGRPCServer(handler *grpcsrv.Handler) *grpc.Server {
	server := grpc.NewServer()
	messenger.RegisterMessengerServiceServer(server, handler)

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
