package linkhandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	errors "github.com/mcandemir/bilinkat/internal/errors"
	linkmodel "github.com/mcandemir/bilinkat/internal/model/link"
	linkservice "github.com/mcandemir/bilinkat/internal/service/link"
)

// LinkHandler definition with LinkService dependency injection
type LinkHandler struct {
	service *linkservice.LinkService
}

// Create a new linkhandler with the given service, inject it
func NewLinkHandler(linkService *linkservice.LinkService) *LinkHandler {
	return &LinkHandler{
		service: linkService,
	}
}

// Shorten handles POST /api/v1/links/shorten
func (h *LinkHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req linkmodel.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteValidationError(w, errors.CodeInvalidInput, "Invalid request body", nil, getRequestID(r))
		return
	}

	// Use service to shorten URL
	link, err := h.service.Shorten(req.URL)
	if err != nil {
		errors.WriteValidationError(w, err.Code, "Invalid URL", nil, getRequestID(r))
		return
	}

	h.sendJSON(w, link, http.StatusCreated)
}

// Redirect handles GET /{slug} - redirects to original URL
func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		errors.WriteValidationError(w, errors.CodeInvalidInput, "Invalid slug", nil, getRequestID(r))
		return
	}

	// Use service to get link
	link, err := h.service.GetLink(slug)
	if err != nil {
		errors.WriteNotFoundError(w, err.Code, err.Message, getRequestID(r))
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
		errors.WriteErrorResponse(w, err, getRequestID(r))
		return
	}

	h.sendJSON(w, links, http.StatusOK)
}

// GetLink handles GET /api/v1/links/{slug}
func (h *LinkHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		errors.WriteValidationError(w, errors.CodeInvalidInput, "Invalid slug", nil, getRequestID(r))
		return
	}

	link, err := h.service.GetLink(slug)

	if err != nil {
		// handle error
		errors.WriteErrorResponse(w, err, getRequestID(r))
		return
	}

	// For now, just use the slug to avoid unused variable error
	h.sendJSON(w, link, http.StatusOK)
}

// UpdateLink handles PUT /api/v1/links/{slug}
func (h *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	var req linkmodel.UpdateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteValidationError(w, errors.CodeInvalidInput, "Invalid request body", nil, getRequestID(r))
		return
	}

	slug := chi.URLParam(r, "slug")
	if slug == "" {
		errors.WriteValidationError(w, errors.CodeInvalidInput, "Invalid slug", nil, getRequestID(r))
		return
	}

	// update service

	link, err := h.service.UpdateLink(slug, &req)
	if err != nil {
		errors.WriteErrorResponse(w, err, getRequestID(r))
		return
	}

	h.sendJSON(w, link, http.StatusOK)
}

// DeleteLink handles DELETE /api/v1/links/{slug}
func (h *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		errors.WriteValidationError(w, errors.CodeInvalidInput, "Invalid slug", nil, getRequestID(r))
		return
	}

	err := h.service.DeleteLink(slug)
	if err != nil {
		errors.WriteErrorResponse(w, err, getRequestID(r))
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

func getRequestID(r *http.Request) string {
	return "123123"
	//return r.Header.Get("X-Request-ID")
}
