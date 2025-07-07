package repository

import (
	"context"

	linkmodel "github.com/mcandemir/bilinkat/internal/model/link"
)

type LinkRepository interface {
	Create(ctx context.Context, link *linkmodel.Link) error
	GetBySlug(ctx context.Context, slug string) (*linkmodel.Link, error)
	GetByUserID(ctx context.Context, userID int) ([]*linkmodel.Link, error)
	Update(ctx context.Context, slug string, link *linkmodel.Link) error
	Delete(ctx context.Context, slug string) error
	SlugExists(ctx context.Context, slug string) (bool, error)
}
