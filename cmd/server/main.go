package main

import (
	"log"
	"log/slog"
	"os"

	config "github.com/mcandemir/bilinkat/internal/config"
	linkhandler "github.com/mcandemir/bilinkat/internal/handler/link"
	logger "github.com/mcandemir/bilinkat/internal/logger"
	router "github.com/mcandemir/bilinkat/internal/router"
	linkservice "github.com/mcandemir/bilinkat/internal/service/link"
)

func main() {
	// load env
	cfg := config.MustLoad()

	// logger
	handlerOpts := logger.NewHandlerOptions(slog.LevelInfo)
	logger := logger.NewLogger("json", handlerOpts, os.Stdout)

	// Initialize service layer
	linkService := linkservice.NewLinkService(cfg, logger)

	// Initialize handler with service dependency
	handlers := &router.Handlers{
		Link: linkhandler.NewLinkHandler(linkService, logger),
	}

	// create routers with handlers dependency
	r := router.NewRouter(handlers, logger, cfg)

	server := NewServer(cfg, r)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
