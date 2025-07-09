# Bİ LİNK AT!

A URL shortening service built with Go, following clean architecture principles.

##### (This project is currently ongoing, using AI supported README and CHANGELOGS to summerize current project status.)

## Features

- ✅ URL shortening with custom slug generation
- ✅ RESTful API with proper error handling
- ✅ Health check endpoint
- ✅ Link management (CRUD operations)
- ✅ URL validation and sanitization
- ✅ Structured logging with request tracking
- ✅ Middleware support (logging, recovery, request ID)
- ✅ Configuration management with environment variables
- ✅ Clean architecture with separation of concerns
- ✅ Comprehensive error handling with custom error types

## Project Structure

```
bilinkat/
├── cmd/
│ └── server/
│ ├── main.go # Application entry point
│ └── server.go # Server configuration
├── internal/
│ ├── config/
│ │ ├── config.go # Configuration structures
│ │ └── loader.go # Environment configuration loader
│ ├── errors/
│ │ ├── codes.go # Error code definitions
│ │ ├── generic.go # Generic error types
│ │ ├── link/
│ │ │ └── link_errors.go # Link-specific errors
│ │ └── response.go # Error response utilities
│ ├── handler/
│ │ └── link/
│ │ └── link_handler.go # HTTP handlers
│ ├── logger/
│ │ └── logger.go # Structured logging
│ ├── middleware/
│ │ └── middlewares.go # HTTP middleware
│ ├── model/
│ │ └── link/
│ │ └── link.go # Data models and DTOs
│ ├── repository/
│ │ └── link/
│ │ └── link_repository.go # Data access layer (placeholder)
│ ├── router/
│ │ └── router.go # Route definitions
│ ├── service/
│ │ └── link/
│ │ └── link_service.go # Business logic
│ ├── utils/
│ │ └── slug_generator.go # Utility functions
│ └── validator/
│ └── link/
│ └── link_validator.go # Input validation
├── go.mod # Go module definition
├── go.sum # Dependency checksums
└── README.md
```

## Architecture

The application follows clean architecture principles with clear separation of concerns:

```
HTTP Request → Router → Handler → Service → Repository → Database
     ↑           ↑         ↑         ↑          ↑
   Middleware  Routing  HTTP Logic  Business   Data Access
                              & Validation    Logic
```

- **Router**: Defines routes and applies middleware using Chi router
- **Handler**: Handles HTTP requests/responses, input validation, and error responses
- **Service**: Contains business logic, URL validation, and slug generation
- **Repository**: Data access layer (currently placeholder for database integration)
- **Model**: Data structures and DTOs for requests/responses
- **Validator**: Input validation for URLs, slugs, and user IDs
- **Middleware**: Cross-cutting concerns (logging, recovery, request ID, real IP)
- **Config**: Environment-based configuration management
- **Errors**: Structured error handling with custom error types

## API Endpoints

### Health Check
```
GET /health
```
Returns service status and version information.

**Response:**
```json
{
  "status": "ok",
  "service": "bilinkat",
  "version": "1.0.0"
}
```

### URL Shortening
```
POST /api/v1/links/shorten
Content-Type: application/json

{
  "url": "https://example.com"
}
```

**Response:**
```json
{
  "id": 1,
  "slug": "abc123",
  "url": "https://example.com",
  "created_at": "2024-01-01T12:00:00Z"
}
```

### Link Management
```
GET    /api/v1/links          # Get user's links
GET    /api/v1/links/{slug}   # Get specific link
PUT    /api/v1/links/{slug}   # Update link
DELETE /api/v1/links/{slug}   # Delete link
```

**Update Link Request:**
```json
{
  "url": "https://updated-example.com"
}
```

### Redirect (Not yet implemented in router)
```
GET /{slug}
```
Redirects to the original URL.

## Configuration

The application uses environment variables for configuration:

### Server Configuration
- `SERVER_PORT` - Server port (default: 3000)
- `SERVER_HOST` - Server host (default: localhost)
- `SERVER_READ_TIMEOUT` - Read timeout (default: 30s)
- `SERVER_WRITE_TIMEOUT` - Write timeout (default: 30s)

### Application Configuration
- `APP_BASE_URL` - Base URL for short links (default: http://localhost:3000)
- `APP_ENV` - Environment (default: development)
- `APP_LOG_LEVEL` - Log level (default: info)

### Database Configuration
- `DATABASE_URL` - Database connection string (required for production)

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd bilinkat
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up environment variables (optional):**
   ```bash
   export SERVER_PORT=3000
   export APP_BASE_URL=http://localhost:3000
   ```

4. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

5. **Test the health check:**
   ```bash
   curl http://localhost:3000/health
   ```

6. **Shorten a URL:**
   ```bash
   curl -X POST http://localhost:3000/api/v1/links/shorten \
     -H "Content-Type: application/json" \
     -d '{"url":"https://example.com"}'
   ```

7. **Get user links:**
   ```bash
   curl http://localhost:3000/api/v1/links
   ```

## Dependencies

- **Chi Router**: Lightweight HTTP router for Go
- **Godotenv**: Environment variable loading
- **Envconfig**: Environment-based configuration
- **Standard Library**: Logging, HTTP, JSON, etc.

## Development

### Code Organization
- **Clean Architecture**: Clear separation between layers
- **Dependency Injection**: Services injected into handlers
- **Error Handling**: Structured error responses with request IDs
- **Validation**: Comprehensive input validation
- **Logging**: Structured logging with request context

### Current Implementation Status
- ✅ HTTP server with Chi router
- ✅ URL shortening service (in-memory)
- ✅ Link management endpoints
- ✅ Input validation
- ✅ Error handling
- ✅ Configuration management
- ✅ Middleware (logging, recovery, request ID)
- ✅ Structured logging
- ⏳ Database integration (repository layer placeholder)
- ⏳ Authentication and authorization
- ⏳ Redirect endpoint implementation
- ⏳ Link analytics
- ⏳ Rate limiting

## Next Steps

- [ ] Implement database integration (PostgreSQL/MySQL)
- [ ] Add user authentication and authorization
- [ ] Implement redirect endpoint in router
- [ ] Add link analytics and click tracking
- [ ] Add rate limiting middleware
- [ ] Add API documentation (Swagger/OpenAPI)
- [ ] Add comprehensive test coverage
- [ ] Add Docker support
- [ ] Add CI/CD pipeline
- [ ] Add monitoring and metrics

## Error Handling

The application uses structured error handling with:
- Custom error types for different scenarios
- HTTP status codes mapping
- Request ID tracking
- Detailed error messages
- Validation error details

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[Add your license here]
