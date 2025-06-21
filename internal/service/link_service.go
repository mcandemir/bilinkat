package service

import (
	"time"

	"github.com/mcandemir/bilinkat/internal/model"
)

type LinkService struct{}

func NewLinkService() *LinkService {
	return &LinkService{}
}

func (s *LinkService) Shorten(url string) (model.Link, error) {
	slug := "testSlug"

	link := model.Link{
		Slug:      slug,
		Url:       url,
		CreatedAt: time.Now(),
	}

	return link, nil
}
