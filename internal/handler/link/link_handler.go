package handler

import (
	"net/http"

	service "github.com/mcandemir/bilinkat/internal/service/link"
)

type LinkHandler struct {
	service *service.LinkService
}

func NewLinkHandler() *LinkHandler {
	return &LinkHandler{}
}

func (h *LinkHandler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	slug, err := h.service.Shorten(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(slug.Slug))
}
