# Go Hexagonal Architecture Project

This project implements a modular monolith REST API in Go, following strict Hexagonal Architecture (Ports & Adapters) principles.

## Features

- **Hexagonal Architecture**: Clear separation of `Domain`, `Application`, `Interfaces`, and `Adapters`.
- **Modular Structure**: Code organized by business modules.
- **Fiber**: High-performance web framework.
- **PostgreSQL & sqlx**: Robust database interactions.
- **Docker & Docker Compose**: Easy deployment and development setup.
- **Auto-Migrations**: Database schema management on startup.

## Prerequisites

- Go 1.22+
- Docker & Docker Compose

## Quick Start

1.  **Clone the repository** (if not already done).

2.  **Run with Docker Compose**:

    ```bash
    docker-compose up --build
    ```

    This will start the PostgreSQL database and the Go application. The application waits for the database to be ready and runs migrations automatically.

3.  **Access the API**:
    The server listens on port `8080`.

## API Endpoints

### Examples Module

- `POST /examples`: Create a new example.
  ```json
  {
    "name": "My Example",
    "description": "This is a test."
  }
  ```
- `GET /examples`: List all examples.
- `GET /examples/:id`: Get a specific example by ID.
- `PUT /examples/:id`: Update an example.
- `DELETE /examples/:id`: Delete an example.

## Project Structure

```
.
├── cmd/                # Application entrypoint
├── internal/
│   ├── bootstrap/      # App initialization and wiring
│   ├── infrastructure/ # Shared infrastructure (DB, Env, Migrations)
│   └── modules/        # Business logic modules
│       └── example/
│           ├── domain/       # Entities & Repository Interfaces
│           ├── application/  # Use Cases
│           ├── interfaces/   # HTTP Handlers (Inbound Adapters)
│           └── adapters/     # Database Repositories (Outbound Adapters)
├── migrations/         # SQL Migration files
├── Dockerfile          # Docker build instructions
└── docker-compose.yml  # Docker services orchestration
```
