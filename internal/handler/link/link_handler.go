package linkhandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	linkmodel "github.com/mcandemir/bilinkat/internal/model/link"
	linkservice "github.com/mcandemir/bilinkat/internal/service/link"
)

type LinkHandler struct {
	service *linkservice.LinkService
}

func NewLinkHandler(linkService *linkservice.LinkService) *LinkHandler {
	return &LinkHandler{
		service: linkService,
	}
}

// Shorten handles POST /api/v1/links/shorten
func (h *LinkHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req linkmodel.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Use service to shorten URL
	link, err := h.service.Shorten(req.URL)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Build response
	baseURL := "http://localhost:3000"
	shortURL := baseURL + "/" + link.Slug

	response := linkmodel.ShortenResponse{
		ShortURL:    shortURL,
		OriginalURL: link.Url,
		Slug:        link.Slug,
	}

	h.sendJSON(w, response, http.StatusCreated)
}

// Redirect handles GET /{slug} - redirects to original URL
func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		h.sendError(w, "Slug is required", http.StatusBadRequest)
		return
	}

	// Use service to get link
	link, err := h.service.GetLink(slug)
	if err != nil {
		h.sendError(w, "Link not found", http.StatusNotFound)
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
}

// GetUserLinks handles GET /api/v1/links
func (h *LinkHandler) GetUserLinks(w http.ResponseWriter, r *http.Request) {
	// For now, use user ID 1 (we'll add authentication later)
	userID := 1

	links, err := h.service.GetUserLinks(userID)
	if err != nil {
		h.sendError(w, "Failed to get links", http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, links, http.StatusOK)
}

// GetLink handles GET /api/v1/links/{slug}
func (h *LinkHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		h.sendError(w, "Link slug is required", http.StatusBadRequest)
		return
	}

	link, err := h.service.GetLink(slug)

	if err != nil {
		// handle error
		log.Default().Println("something wrong with get link")
	}

	// For now, just use the slug to avoid unused variable error
	h.sendJSON(w, link, 200)
}

// UpdateLink handles PUT /api/v1/links/{slug}
func (h *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "slug")
	if idStr == "" {
		h.sendError(w, "Link slug is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, "Invalid link slug", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	link, err := h.service.UpdateLink(id, updates)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusNotFound)
		return
	}

	h.sendJSON(w, link, http.StatusOK)
}

// DeleteLink handles DELETE /api/v1/links/{id}
func (h *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		h.sendError(w, "Link ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteLink(id)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper methods
func (h *LinkHandler) sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *LinkHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	response := linkmodel.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	h.sendJSON(w, response, statusCode)
}
