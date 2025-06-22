package main

import (
	"net/http"

	handler "github.com/mcandemir/bilinkat/internal/handler/link"
	router "github.com/mcandemir/bilinkat/internal/router"
)

func main() {
	handlers := &router.Handlers{
		Link: handler.NewLinkHandler(),
	}

	router := router.NewRouter(*handlers)

	http.ListenAndServe(":3000", router)
}
