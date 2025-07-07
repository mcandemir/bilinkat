package linkservice

import (
	"time"

	"context"

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
func (s *LinkService) Shorten(ctx context.Context, url string) (*linkmodel.Link, *errors.AppError) {
	s.logger.Debug(ctx, "Starting URL shortening process", "url", url)

	if err := linkvalidator.NewLinkValidator().ValidateURL(url); err != nil {
		s.logger.Error(ctx, "URL validation failed",
			"url", url,
			"error_code", err.Code,
			"error_message", err.Message,
			"validation_details", err.Details)
		return nil, err
	}

	slug, err := s.generateUniqueSlug()
	if err != nil {
		s.logger.Error(ctx, "Failed to generate unique slug",
			"error_code", err.Code,
			"error_message", err.Message)
		return nil, err
	}

	link := &linkmodel.Link{
		Slug:      slug,
		Url:       url,
		CreatedAt: time.Now(),
	}

	s.logger.Info(ctx, "URL shortened successfully",
		"original_url", url,
		"slug", slug)

	return link, nil
}

// GetLink retrieves a link by its slug
func (s *LinkService) GetLink(ctx context.Context, slug string) (*linkmodel.Link, *errors.AppError) {
	s.logger.Debug(ctx, "Starting link retrieval process", "slug", slug)

	err := linkvalidator.NewLinkValidator().ValidateSlug(slug)
	if err != nil {
		s.logger.Error(ctx, "Invalid slug", "slug", slug, "error_code", err.Code, "error_message", err.Message)
		return nil, err
	}

	// TODO: Get from database via repository
	// For now, return a mock link
	link := linkmodel.CreateExampleLink(slug)

	return link, nil
}

// GetUserLinks retrieves all links for a user
func (s *LinkService) GetUserLinks(ctx context.Context, userID int) ([]*linkmodel.Link, *errors.AppError) {
	s.logger.Debug(ctx, "Starting user links retrieval process", "user_id", userID)

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
func (s *LinkService) UpdateLink(ctx context.Context, slug string, updateLinkModel *linkmodel.UpdateLinkRequest) (*linkmodel.Link, *errors.AppError) {
	s.logger.Debug(ctx, "Starting link update process", "slug", slug)

	if err := linkvalidator.NewLinkValidator().ValidateURL(updateLinkModel.URL); err != nil {
		s.logger.Error(ctx, "Invalid URL", "url", updateLinkModel.URL, "error_code", err.Code, "error_message", err.Message)
		return nil, err
	}

	// if valid, conduct the updating process
	link := linkmodel.CreateExampleLink(slug)

	return link, nil
}

// DeleteLink deletes a link
func (s *LinkService) DeleteLink(ctx context.Context, slug string) *errors.AppError {
	s.logger.Debug(ctx, "Starting link deletion process", "slug", slug)

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
