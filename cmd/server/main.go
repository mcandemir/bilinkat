package main

import (
	"net/http"

	linkhandler "github.com/mcandemir/bilinkat/internal/handler/link"
	router "github.com/mcandemir/bilinkat/internal/router"
	linkservice "github.com/mcandemir/bilinkat/internal/service/link"
)

func main() {
	// Initialize service layer
	linkService := linkservice.NewLinkService()

	// Initialize handler with service dependency
	handlers := &router.Handlers{
		Link: linkhandler.NewLinkHandler(linkService),
	}

	router := router.NewRouter(handlers)

	http.ListenAndServe(":3000", router)
}
