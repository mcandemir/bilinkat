package service

import (
	"time"

	model "github.com/mcandemir/bilinkat/internal/model/link"
	"github.com/mcandemir/bilinkat/internal/utils"
)

type LinkService struct{}

func NewLinkService() *LinkService {
	return &LinkService{}
}

func (s *LinkService) Shorten(url string) (model.Link, error) {
	slug := utils.GenerateSlug(7)

	link := model.Link{
		Slug:      slug,
		Url:       url,
		CreatedAt: time.Now(),
	}

	return link, nil
}
