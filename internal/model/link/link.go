package linkmodel

import "time"

type Link struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateExampleLink(slug string) *Link {
	return &Link{
		Id:        1,
		Slug:      slug,
		Url:       "https://mehmetcandemir.com",
		CreatedAt: time.Now(),
	}
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	Slug        string `json:"slug"`
}

type UpdateLinkRequest struct {
	URL string `json:"url"`
}

type UpdateLinkResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	Slug        string `json:"slug"`
}
