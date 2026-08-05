## About

This project is a REST API for a blogging platform built with Go. It demonstrates a layered architecture, JWT authentication, PostgreSQL integration, request validation, and structured logging.

# Blog API

A RESTful blogging platform built with Go.

## Features

- User registration and authentication (JWT)
- CRUD operations for blog posts
- Request validation
- PostgreSQL with GORM
- Structured logging
- Configuration using Viper
- Database migrations

## Tech Stack

- Go
- Gin
- PostgreSQL
- GORM
- JWT
- Viper
- Zap Logger

## Project Structure

```
cmd/            # application entry point
api/            # HTTP handlers, routes and middleware
config/         # application configuration
data/           # database models and migrations
services/       # business logic
pkg/            # reusable packages
```

## Getting Started

```bash
go mod tidy
go run ./cmd
```

## Roadmap

- [ ] Docker Compose
- [ ] Swagger documentation
- [ ] Refresh tokens
- [ ] Unit tests
- [ ] Pagination
- [ ] Rate limiting