package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/mcandemir/bilinkat/internal/config"
	router "github.com/mcandemir/bilinkat/internal/router"
)

type Server struct {
	config *config.Config
	router *router.Router
	server *http.Server
}

func NewServer(cfg *config.Config, r *router.Router) *Server {
	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port

	return &Server{
		config: cfg,
		router: r,
		server: &http.Server{
			Addr:              serverAddr,
			Handler:           r,
			ReadHeaderTimeout: cfg.Server.ReadTimeout,
			WriteTimeout:      cfg.Server.WriteTimeout,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s", s.server.Addr)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited")
	return nil
}
