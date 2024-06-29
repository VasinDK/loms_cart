package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Config interface {
	GetPort() string
	GetTimeGraceShutdown() int
}

type Server struct {
	HttpServer        http.Server
	TimeGraceShutdown int
}

// New - создает http сервер
func New(mux http.Handler, config Config) *Server {
	return &Server{
		HttpServer: http.Server{
			Handler: mux,
			Addr:    fmt.Sprintf(":%v", config.GetPort()),
		},
		TimeGraceShutdown: config.GetTimeGraceShutdown(),
	}
}

// Run - запускает http сервер
func (s *Server) Run() <-chan error {
	chLis := make(chan error)

	go func() {
		err := s.HttpServer.ListenAndServe()
		if err != nil {
			chLis <- err
		}
	}()

	return chLis
}

// GraceShutdown - плавно завершает сервер
func (s *Server) GraceShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.TimeGraceShutdown)*time.Second)
	defer cancel()

	err := s.HttpServer.Shutdown(ctx)
	if err != nil {
		slog.Info("s.HttpServer.Shutdown", "err", err)
	}
}
