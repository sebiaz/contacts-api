# Contacts API

A simple REST API for managing contacts built with Go, Gin, GORM, and SQLite.

## Features

- CRUD operations for contacts
- Search contacts by name, email, or company
- SQLite database (file-based, no setup required)
- CORS enabled for frontend integration
- Input validation

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/contacts` | List all contacts (supports `?q=search`) |
| GET | `/api/v1/contacts/:id` | Get a single contact |
| POST | `/api/v1/contacts` | Create a new contact |
| PUT | `/api/v1/contacts/:id` | Update a contact |
| DELETE | `/api/v1/contacts/:id` | Delete a contact |

## Contact Fields

- `first_name` (required)
- `last_name` (required)
- `email` (required, unique)
- `phone`
- `address`
- `company`
- `job_title`
- `notes`

## Quick Start

### Prerequisites

- Go 1.23+

### Run the API

```bash
# Download dependencies
go mod download

# Run the server
go run main.go
```

The server will start on `http://localhost:8080`

### Example Requests

**Create a contact:**
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

**List all contacts:**
```bash
curl http://localhost:8080/api/v1/contacts
```

**Search contacts:**
```bash
curl "http://localhost:8080/api/v1/contacts?q=john"
```

**Get a single contact:**
```bash
curl http://localhost:8080/api/v1/contacts/1
```

**Update a contact:**
```bash
curl -X PUT http://localhost:8080/api/v1/contacts/1 \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1-555-9999"}'
```

**Delete a contact:**
```bash
curl -X DELETE http://localhost:8080/api/v1/contacts/1
```

## Project Structure

```
contacts-api/
├── go.mod
├── main.go
├── contacts.db          # Created automatically
└── internal/
    ├── models/
    │   └── contact.go
    ├── database/
    │   └── contact_db.go
    └── handlers/
        └── contact_handler.go
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver