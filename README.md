# Contacts API

A RESTful API for managing contacts built with Go, Gin, GORM, and SQLite.

## Features

- CRUD operations for contacts (Create, Read, Update, Delete)
- Search contacts by name, email, phone, or company
- Pagination support
- SQLite database (file-based, no external setup required)
- CORS enabled for frontend integration
- Input validation with Gin binding tags
- Structured project layout (models, database, handlers)

## Tech Stack

- **Language**: Go 1.23+
- **Framework**: Gin (HTTP web framework)
- **ORM**: GORM
- **Database**: SQLite via modernc.org/sqlite (pure Go driver)
- **Validation**: go-playground/validator/v10 (via Gin)

## Project Structure

```
contacts-api/
├── main.go                      # Application entry point
├── go.mod / go.sum              # Go module dependencies
├── contacts.db                  # SQLite database (gitignored)
├── .env.example                 # Environment variable template
├── .gitignore
├── README.md
├── contact.json                 # Sample create request
├── update.json                  # Sample update request
└── internal/
    ├── models/
    │   └── contact.go           # Contact model & DTOs
    ├── database/
    │   └── contact_db.go        # Database interface & implementation
    └── handlers/
        └── contact_handler.go   # HTTP handlers
```

## Quick Start

### Prerequisites

- Go 1.23 or later

### Installation

```bash
# Clone the repository
git clone https://github.com/sebiaz/contacts-api.git
cd contacts-api

# Copy environment template and adjust if needed
cp .env.example .env

# Download dependencies
go mod download

# Run the server
go run main.go
```

The server starts on `http://localhost:8080` (or the `PORT` environment variable).

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port |
| `DB_PATH` | `contacts.db` | Path to SQLite database file |
| `GIN_MODE` | `debug` | Gin mode: `debug`, `release`, `test` |
| `CORS_ALLOWED_ORIGINS` | `*` | Comma-separated allowed origins for CORS |

See `.env.example` for all available options.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/api/v1/contacts` | List contacts (supports `?q=search&page=1&page_size=20`) |
| `GET` | `/api/v1/contacts/:id` | Get single contact |
| `POST` | `/api/v1/contacts` | Create contact |
| `PUT` | `/api/v1/contacts/:id` | Update contact |
| `DELETE` | `/api/v1/contacts/:id` | Delete contact |

### Contact Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `first_name` | string | Yes | Contact's first name |
| `last_name` | string | Yes | Contact's last name |
| `email` | string | Yes | Unique email address |
| `phone` | string | No | Phone number |
| `address` | string | No | Physical address |
| `company` | string | No | Company name |
| `job_title` | string | No | Job title |
| `notes` | string | No | Additional notes |

## Example Requests

### Create a Contact

```bash
curl -X POST http://localhost:8080/api/v1/contacts \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "phone": "+1-555-0123",
    "company": "Acme Corp",
    "job_title": "Software Engineer"
  }'
```

### List All Contacts

```bash
curl http://localhost:8080/api/v1/contacts
```

### Search Contacts

```bash
curl "http://localhost:8080/api/v1/contacts?q=john"
```

### Get Single Contact

```bash
curl http://localhost:8080/api/v1/contacts/1
```

### Update a Contact

```bash
curl -X PUT http://localhost:8080/api/v1/contacts/1 \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1-555-9999", "notes": "Updated via API"}'
```

### Delete a Contact

```bash
curl -X DELETE http://localhost:8080/api/v1/contacts/1
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o contacts-api main.go
```

### Linting

```bash
golangci-lint run
```

## Security Considerations

- **CORS**: Default allows all origins (`*`). Configure `CORS_ALLOWED_ORIGINS` for production.
- **Authentication**: Not implemented. Add auth middleware for production use.
- **Rate Limiting**: Not implemented. Consider adding for public APIs.
- **Input Validation**: Uses Gin binding tags for basic validation.
- **SQL Injection**: Prevented by using GORM parameterized queries.
- **Dependencies**: Regularly run `govulncheck ./...` to check for vulnerabilities.

## License

MIT License