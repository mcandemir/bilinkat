package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	linkhandler "github.com/mcandemir/bilinkat/internal/handler/link"
	"github.com/mcandemir/bilinkat/internal/middleware"
)

type Handlers struct {
	Link *linkhandler.LinkHandler
}

type Router struct {
	chi.Router
	handlers *Handlers
}

func NewRouter(handlers *Handlers) *Router {
	router := &Router{
		Router:   chi.NewRouter(),
		handlers: handlers,
	}

	router.setupMiddleware()
	router.setupRoutes()

	return router
}

func (r *Router) setupMiddleware() {
	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
}

func (r *Router) setupRoutes() {
	// Health check
	r.Get("/health", r.healthCheck)

	// API routes
	r.Route("/api/v1", func(v1 chi.Router) {
		r.setupLinkRoutes(v1)
	})
}

func (r *Router) setupLinkRoutes(v1 chi.Router) {
	v1.Route("/links", func(links chi.Router) {
		links.Post("/shorten", r.handlers.Link.Shorten)
		links.Get("/", r.handlers.Link.GetUserLinks)
		links.Get("/{id}", r.handlers.Link.GetLink)
		links.Put("/{id}", r.handlers.Link.UpdateLink)
		links.Delete("/{id}", r.handlers.Link.DeleteLink)
	})
}

func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","service":"bilinkat","version":"1.0.0"}`))
}
