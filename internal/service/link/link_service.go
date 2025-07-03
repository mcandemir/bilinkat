package linkservice

import (
	"time"

	"github.com/mcandemir/bilinkat/internal/config"
	errors "github.com/mcandemir/bilinkat/internal/errors"
	logger "github.com/mcandemir/bilinkat/internal/logger"
	linkmodel "github.com/mcandemir/bilinkat/internal/model/link"
	utils "github.com/mcandemir/bilinkat/internal/utils"
	linkvalidator "github.com/mcandemir/bilinkat/internal/validator/link"
)

type LinkService struct {
	config *config.Config
	logger *logger.Logger
}

func NewLinkService(cfg *config.Config, logger *logger.Logger) *LinkService {
	return &LinkService{
		config: cfg,
		logger: logger,
	}
}

// Shorten creates a shortened URL for the given long URL
func (s *LinkService) Shorten(url string) (*linkmodel.Link, *errors.AppError) {
	// Validate URL
	if err := linkvalidator.NewLinkValidator().ValidateURL(url); err != nil {
		return nil, err
	}

	// Generate unique slug
	slug, err := s.generateUniqueSlug()
	if err != nil {
		return nil, err
	}

	// Create link
	link := &linkmodel.Link{
		Slug:      slug,
		Url:       url,
		CreatedAt: time.Now(),
	}

	// TODO: Save to database via repository
	// For now, just return the link

	return link, nil
}

// GetLink retrieves a link by its slug
func (s *LinkService) GetLink(slug string) (*linkmodel.Link, *errors.AppError) {
	err := linkvalidator.NewLinkValidator().ValidateSlug(slug)
	if err != nil {
		return nil, err
	}

	// TODO: Get from database via repository
	// For now, return a mock link
	link := linkmodel.CreateExampleLink(slug)

	return link, nil
}

// GetUserLinks retrieves all links for a user
func (s *LinkService) GetUserLinks(userID int) ([]*linkmodel.Link, *errors.AppError) {
	// TODO: Get from database via repository
	// For now, return empty list
	links := []*linkmodel.Link{}
	links = append(links, linkmodel.CreateExampleLink("123456"))
	links = append(links, linkmodel.CreateExampleLink("123457"))
	links = append(links, linkmodel.CreateExampleLink("123458"))
	links = append(links, linkmodel.CreateExampleLink("123459"))
	links = append(links, linkmodel.CreateExampleLink("123460"))
	return links, nil
}

// UpdateLink updates an existing link
func (s *LinkService) UpdateLink(slug string, updateLinkModel *linkmodel.UpdateLinkRequest) (*linkmodel.Link, *errors.AppError) {
	// validate request context
	if err := linkvalidator.NewLinkValidator().ValidateURL(updateLinkModel.URL); err != nil {
		return nil, err
	}

	// if valid, conduct the updating process
	link := linkmodel.CreateExampleLink(slug)

	return link, nil
}

// DeleteLink deletes a link
func (s *LinkService) DeleteLink(slug string) *errors.AppError {
	// TODO: Delete from database via repository
	// For now, return success
	return nil
}

// generateUniqueSlug creates a unique slug
func (s *LinkService) generateUniqueSlug() (string, *errors.AppError) {
	const maxAttempts = 10
	const slugLength = 6

	for attempt := 0; attempt < maxAttempts; attempt++ {
		slug := utils.GenerateSlug(slugLength)

		// TODO: Check if slug exists in database
		// For now, assume it's unique
		return slug, nil
	}

	return "", errors.NewInternalError("failed to generate unique slug after maximum attempts", errors.CodeExternalServiceError, nil)
}
