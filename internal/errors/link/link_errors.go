package linkerrors

import errors "github.com/mcandemir/bilinkat/internal/errors"

const (
	CodeInvalidURL    = "INVALID_URL"
	CodeInvalidSlug   = "INVALID_SLUG"
	CodeSlugExists    = "SLUG_EXISTS"
	CodeLinkNotFound  = "LINK_NOT_FOUND"
	CodeLinkExpired   = "LINK_EXPIRED"
	CodeLinkDisabled  = "LINK_DISABLED"
	CodeUserNotFound  = "USER_NOT_FOUND"
	CodeQuotaExceeded = "QUOTA_EXCEEDED"
)

func NewInvalidURLError(message string, details map[string]interface{}) *errors.AppError {
	return errors.NewValidationError(CodeInvalidURL, message, details)
}

func NewInvalidSlugError(message string, details map[string]interface{}) *errors.AppError {
	return errors.NewValidationError(CodeInvalidSlug, message, details)
}

func NewSlugExistsError(slug string) *errors.AppError {
	return errors.NewConflictError(CodeSlugExists, "Slug already exists")
}

func NewLinkNotFoundError(slug string) *errors.AppError {
	return errors.NewNotFoundError(CodeLinkNotFound, "Link not found")
}

func NewLinkExpiredError(slug string) *errors.AppError {
	return errors.NewValidationError(CodeLinkExpired, "Link has expired", nil)
}

func NewLinkDisabledError(slug string) *errors.AppError {
	return errors.NewValidationError(CodeLinkDisabled, "Link is disabled", nil)
}

func NewUserNotFoundError(userID int) *errors.AppError {
	return errors.NewNotFoundError(CodeUserNotFound, "User not found")
}

func NewQuotaExceededError(message string) *errors.AppError {
	return errors.NewValidationError(CodeQuotaExceeded, message, nil)
}
