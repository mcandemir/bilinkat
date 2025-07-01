package errors

// Common error codes used across the application
const (
	// Validation errors
	CodeInvalidInput  = "INVALID_INPUT"
	CodeMissingField  = "MISSING_FIELD"
	CodeInvalidFormat = "INVALID_FORMAT"
	CodeInvalidLength = "INVALID_LENGTH"
	CodeInvalidValue  = "INVALID_VALUE"

	// Resource errors
	CodeResourceNotFound = "RESOURCE_NOT_FOUND"
	CodeResourceExists   = "RESOURCE_EXISTS"
	CodeResourceDeleted  = "RESOURCE_DELETED"

	// Authentication errors
	CodeUnauthorized            = "UNAUTHORIZED"
	CodeInvalidToken            = "INVALID_TOKEN"
	CodeTokenExpired            = "TOKEN_EXPIRED"
	CodeInsufficientPermissions = "INSUFFICIENT_PERMISSIONS"

	// Database errors
	CodeDatabaseError    = "DATABASE_ERROR"
	CodeConnectionError  = "CONNECTION_ERROR"
	CodeTransactionError = "TRANSACTION_ERROR"

	// External service errors
	CodeExternalServiceError = "EXTERNAL_SERVICE_ERROR"
	CodeTimeoutError         = "TIMEOUT_ERROR"

	// Business logic errors
	CodeBusinessRuleViolation = "BUSINESS_RULE_VIOLATION"
	CodeRateLimitExceeded     = "RATE_LIMIT_EXCEEDED"
	CodeQuotaExceeded         = "QUOTA_EXCEEDED"
)
