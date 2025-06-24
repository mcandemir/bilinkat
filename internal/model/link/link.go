package linkmodel

import "time"

type Link struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	Slug        string `json:"slug"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
