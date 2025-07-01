package main

import (
	"log"

	config "github.com/mcandemir/bilinkat/internal/config"
	linkhandler "github.com/mcandemir/bilinkat/internal/handler/link"
	router "github.com/mcandemir/bilinkat/internal/router"
	linkservice "github.com/mcandemir/bilinkat/internal/service/link"
)

func main() {
	// load env
	cfg := config.MustLoad()

	// Initialize service layer
	linkService := linkservice.NewLinkService(cfg)

	// Initialize handler with service dependency
	handlers := &router.Handlers{
		Link: linkhandler.NewLinkHandler(linkService),
	}

	// create routers with handlers dependency
	r := router.NewRouter(handlers)

	server := NewServer(cfg, r)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
