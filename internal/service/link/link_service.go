package linkservice

import (
	"errors"
	"fmt"
	"strings"
	"time"

	linkmodel "github.com/mcandemir/bilinkat/internal/model/link"
	"github.com/mcandemir/bilinkat/internal/utils"
)

type LinkService struct{}

func NewLinkService() *LinkService {
	return &LinkService{}
}

// Shorten creates a shortened URL for the given long URL
func (s *LinkService) Shorten(url string) (*linkmodel.Link, error) {
	// Validate URL
	if err := s.validateURL(url); err != nil {
		return nil, err
	}

	// Generate unique slug
	slug, err := s.generateUniqueSlug()
	if err != nil {
		return nil, fmt.Errorf("failed to generate slug: %w", err)
	}

	// Create link
	link := &linkmodel.Link{
		Slug:      slug,
		Url:       url,
		CreatedAt: time.Now(),
	}

	// TODO: Save to database via repository
	// For now, just return the link
	fmt.Printf("Created link: %s -> %s\n", slug, url)

	return link, nil
}

// GetLink retrieves a link by its slug
func (s *LinkService) GetLink(slug string) (*linkmodel.Link, error) {
	if slug == "" {
		return nil, errors.New("slug cannot be empty")
	}

	// TODO: Get from database via repository
	// For now, return a mock link
	link := &linkmodel.Link{
		Id:        1,
		Slug:      slug,
		Url:       "https://mehmetcandemir.com",
		CreatedAt: time.Now(),
	}

	return link, nil
}

// GetUserLinks retrieves all links for a user
func (s *LinkService) GetUserLinks(userID int) ([]*linkmodel.Link, error) {
	// TODO: Get from database via repository
	// For now, return empty list
	return []*linkmodel.Link{}, nil
}

// UpdateLink updates an existing link
func (s *LinkService) UpdateLink(id int, updates map[string]interface{}) (*linkmodel.Link, error) {
	// TODO: Update in database via repository
	// For now, return error
	return nil, errors.New("link not found")
}

// DeleteLink deletes a link
func (s *LinkService) DeleteLink(id int) error {
	// TODO: Delete from database via repository
	// For now, return success
	return nil
}

// validateURL checks if the URL is valid
func (s *LinkService) validateURL(url string) error {
	if url == "" {
		return errors.New("URL cannot be empty")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return errors.New("URL must start with http:// or https://")
	}

	return nil
}

// generateUniqueSlug creates a unique slug
func (s *LinkService) generateUniqueSlug() (string, error) {
	const maxAttempts = 10
	const slugLength = 6

	for attempt := 0; attempt < maxAttempts; attempt++ {
		slug := utils.GenerateSlug(slugLength)

		// TODO: Check if slug exists in database
		// For now, assume it's unique
		return slug, nil
	}

	return "", errors.New("failed to generate unique slug after maximum attempts")
}
