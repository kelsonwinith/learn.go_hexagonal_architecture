# AI Coding Agent Rules & Architecture Guidelines

You are required to follow these strict architectural rules.

Do not violate them.

---

# Architecture Style

This project uses:

- Hexagonal Architecture (Ports & Adapters)
- Modular Monolith
- Clean separation of concerns
- Strict inward dependency rule

---

# Folder Structure

internal/
  modules/
    example/
      domain/
      application/
      interfaces/http/
      adapters/postgresql/
      example-module.go
  infrastructure/
    postgresql/
    middleware/
  bootstrap/

---

# Naming Rules

1. All files inside module MUST start with module name.

   example-entity.go
   example-create.go
   example-handler-create.go
   example-postgresql-repository.go

2. Use kebab-case for filenames.
3. 1 file = 1 execute use case.
4. Do not group multiple use cases in one file.

---

# Layer Responsibilities

## Domain Layer

- Contains:
  - Entity
  - Repository interface
  - Domain errors
- Must NOT import:
  - Fiber
  - sqlx
  - database/sql
  - any framework
- Pure business logic only.

---

## Application Layer

- Contains use cases.
- Each use case has:

  type UseCase struct {}
  func (u *UseCase) Execute(...)

- Accept structured input.
- Return structured output.
- Must NOT contain SQL.
- Must NOT contain HTTP logic.
- Orchestrates domain + repository.

---

## Interfaces Layer (Inbound)

Path:
modules/example/interfaces/http/

- Contains HTTP handlers.
- Handles:
  - request parsing
  - validation
  - calling use case
  - response formatting
- Must NOT contain business logic.
- Must NOT access database directly.

---

## Adapters Layer (Outbound)

Path:
modules/example/adapters/postgresql/

- Implements domain repository interface.
- Uses sqlx.
- Contains SQL queries.
- Maps DB rows to domain entity.
- No business rules here.

---

## Infrastructure Layer

internal/infrastructure/

Contains:

- Database connection setup
- Middleware
- Shared technical utilities

Must NOT contain module business logic.

---

# Dependency Rule (Critical)

Allowed direction:

interfaces → application → domain → adapters

Adapters may depend on infrastructure.

Domain must depend on nothing.

Never reverse dependency direction.

---

# Database Rules

- Use sqlx.
- Use context in every query.
- Use prepared statements or named queries.
- Handle SQL errors properly.
- Do not leak raw SQL errors to HTTP layer.
- Map DB models to domain entity cleanly.

---

# Error Handling

- Domain errors defined in domain layer.
- Application wraps errors if needed.
- HTTP layer maps errors to proper HTTP status codes.
- No panic for business logic.

---

# Docker Requirements

- Multi-stage Dockerfile.
- Use official Go image for build.
- Use minimal base image for runtime.
- docker-compose must:
  - Start PostgreSQL
  - Start app
  - Pass env vars
- App must wait for DB before starting.

---

# Migration Rules

- Create SQL migration file.
- Auto-run migration on app startup.
- Do not use external migration tool.
- Keep migration simple.

---

# Code Quality Rules

- Use UUID v4.
- Use time.Now().UTC()
- No global variables except config/bootstrap.
- Proper dependency injection via constructor.
- No circular dependencies.
- No shared folder dumping.

---

# Final Rule

The generated project MUST:

- Compile successfully
- Run successfully via docker compose
- Follow all architectural constraints above
- Respect modular + hexagonal discipline
