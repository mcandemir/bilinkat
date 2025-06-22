package router

import (
	"github.com/go-chi/chi/v5"
	handler "github.com/mcandemir/bilinkat/internal/handler/link"
)

type Handlers struct {
	Link *handler.LinkHandler
}

type Router struct {
	chi.Router
	linkHandler *handler.LinkHandler
}

func NewRouter(handlers Handlers) *Router {
	r := &Router{
		Router:      chi.NewRouter(),
		linkHandler: handlers.Link,
	}
	r.setupApiRoutes()

	return r
}

func (r *Router) setupApiRoutes() {
	r.Router.Route("/api/v1", func(v1 chi.Router) {
		r.setupApiLinkRoutes(v1)
	})
}

func (r *Router) setupApiLinkRoutes(v1 chi.Router) {
	v1.Route("/link", func(link chi.Router) {
		link.Get("/", r.linkHandler.Shorten)
	})
}
