# Bİ LİNK AT!

A URL shortening service built with Go.

## Features

- ✅ URL shortening
- ✅ Health check endpoint
- ✅ RESTful API
- ✅ Proper error handling
- ✅ JSON responses
- ✅ Middleware support
- ✅ Service layer architecture
- ✅ Clean package naming

## Project Structure

```
bilinkat/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── handler/
│   │   └── link/
│   │       └── link_handler.go  # HTTP handlers
│   ├── middleware/
│   │   └── logger.go        # HTTP middleware
│   ├── model/
│   │   └── link/
│   │       └── link.go      # Data models
│   ├── service/
│   │   └── link/
│   │       └── link_service.go  # Business logic
│   ├── utils/
│   │   └── slug_generator.go # Utility functions
│   └── router/
│       └── router.go        # Route definitions
└── README.md
```

## Architecture

The application follows clean architecture principles:

```
HTTP Request → Handler → Service → Repository → Database
     ↑           ↑         ↑          ↑
   Router    HTTP Logic  Business   Data Access
             & Validation  Logic
```

- **Router**: Defines routes and applies middleware
- **Handler**: Handles HTTP requests/responses, validation
- **Service**: Contains business logic, URL validation, slug generation
- **Model**: Data structures
- **Utils**: Helper functions for slug generation

## API Endpoints

### Health Check
```
GET /health
```
Returns service status.

### URL Shortening
```
POST /api/v1/links/shorten
Content-Type: application/json

{
  "url": "https://example.com"
}
```

Response:
```json
{
  "short_url": "http://localhost:3000/abc123",
  "original_url": "https://example.com",
  "slug": "abc123"
}
```

### Redirect
```
GET /{slug}
```
Redirects to the original URL.

### Link Management
```
GET    /api/v1/links          # Get user's links
GET    /api/v1/links/{id}     # Get specific link
PUT    /api/v1/links/{id}     # Update link
DELETE /api/v1/links/{id}     # Delete link
```

## Getting Started

1. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

2. **Test the health check:**
   ```bash
   curl http://localhost:3000/health
   ```

3. **Shorten a URL:**
   ```bash
   curl -X POST http://localhost:3000/api/v1/links/shorten \
     -H "Content-Type: application/json" \
     -d '{"url":"https://example.com"}'
   ```

4. **Test redirect:**
   ```bash
   curl -L http://localhost:3000/abc123
   ```

## Environment Variables

- `PORT` - Server port (default: 3000)

## Development

The code is structured following clean architecture principles:

- **Handlers**: Handle HTTP requests/responses
- **Service**: Business logic, URL validation, slug generation
- **Router**: Define routes and middleware
- **Middleware**: Cross-cutting concerns (logging, recovery, etc.)
- **Utils**: Helper functions for slug generation

## Next Steps

- [ ] Add repository layer for data access
- [ ] Add database integration
- [ ] Add user authentication
- [ ] Add link analytics
- [ ] Add link-in-bio features
