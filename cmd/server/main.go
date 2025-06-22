package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	handler "github.com/mcandemir/bilinkat/internal/handler/link"
)

func main() {
	r := chi.NewRouter()
	linkHandler := handler.NewLinkHandler()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bİ LİNK AT!"))
	})

	r.Get("/shorten", linkHandler.ShortenHandler)

	http.ListenAndServe(":3000", r)
}
