package handler

import (
	"net/http"

	"github.com/mcandemir/bilinkat/internal/service"
)

type LinkHandler struct {
	service *service.LinkService
}

func NewLinkHandler() *LinkHandler {
	return &LinkHandler{}
}

func (h *LinkHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	slug, err := h.service.Shorten(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(slug.Slug))
}
