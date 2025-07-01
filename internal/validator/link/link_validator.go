package link

import (
	"net/url"
	"regexp"
	"strings"

	errors "github.com/mcandemir/bilinkat/internal/errors"
	linkerrors "github.com/mcandemir/bilinkat/internal/errors/link"
)

// LinkValidator contains validation methods for link-related operations
type LinkValidator struct{}

// NewLinkValidator creates a new link validator
func NewLinkValidator() *LinkValidator {
	return &LinkValidator{}
}

// ValidateURL checks if the URL is valid and follows business rules
func (v *LinkValidator) ValidateURL(urlStr string) *errors.AppError {
	// check if url is empty
	if urlStr == "" {
		return linkerrors.NewInvalidURLError("URL cannot be empty", map[string]interface{}{
			"field": "url",
			"value": urlStr,
		})
	}

	// Check protocol
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		return linkerrors.NewInvalidURLError("URL must start with http:// or https://", map[string]interface{}{
			"field": "url",
			"value": urlStr,
		})
	}

	// Parse URL to check if it's well-formed
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return linkerrors.NewInvalidURLError("Failed to parse URL", map[string]interface{}{
			"field": "url",
			"value": urlStr,
		})
	}

	// Check if host is present
	if parsedURL.Host == "" {
		return linkerrors.NewInvalidURLError("URL must contain a valid host", map[string]interface{}{
			"field": "url",
			"value": urlStr,
		})
	}

	return nil
}

// ValidateSlug checks if the slug is valid
func (v *LinkValidator) ValidateSlug(slug string) *errors.AppError {
	if slug == "" {
		return linkerrors.NewInvalidSlugError("Slug cannot be empty", map[string]interface{}{
			"field": "slug",
			"value": slug,
		})
	}

	// Check length
	if len(slug) != 6 {
		return linkerrors.NewInvalidSlugError("Slug must be 6 characters", map[string]interface{}{
			"field": "slug",
			"value": slug,
		})
	}

	// Check format (alphanumeric)
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, slug)
	if !matched {
		return linkerrors.NewInvalidSlugError("Slug can only contain alphanumeric characters", map[string]interface{}{
			"field": "slug",
			"value": slug,
		})
	}

	return nil
}

// ValidateUserID checks if the user ID is valid
func (v *LinkValidator) ValidateUserID(userID int) *errors.AppError {
	if userID <= 0 {
		return linkerrors.NewUserNotFoundError(userID)
	}
	return nil
}
