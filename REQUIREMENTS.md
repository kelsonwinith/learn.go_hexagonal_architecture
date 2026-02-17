# Project Requirements

## Overview

Build a modular monolith REST API in Go using:

- Go 1.22+
- Fiber framework
- sqlx for database access
- PostgreSQL
- Docker & docker-compose

The project must follow a Hexagonal Architecture (Ports & Adapters) with a modular structure.

---

## Functional Requirements

### Module: Example

Create an `examples` table with full CRUD capability.

### Database Schema

Table: examples

Columns:

- id (UUID, primary key)
- name (varchar, not null)
- description (text, nullable)
- created_at (timestamp, not null)
- updated_at (timestamp, not null)

Use PostgreSQL UUID generation.

---

## REST Endpoints

Base path: /examples

1. Create Example
   - POST /examples
   - Request:
     {
     "name": "string",
     "description": "string"
     }
   - Response: created object

2. Get All Examples
   - GET /examples
   - Response: list of examples

3. Get Example By ID
   - GET /examples/:id
   - Response: example object

4. Update Example
   - PUT /examples/:id
   - Request:
     {
     "name": "string",
     "description": "string"
     }
   - Response: updated object

5. Delete Example
   - DELETE /examples/:id
   - Response: 204 No Content

---

## Non-Functional Requirements

- Clean hexagonal architecture
- Modular structure
- 1 file = 1 use case execute
- Clear separation of:
  - domain
  - application
  - interfaces (inbound)
  - adapters (outbound)
- No business logic inside HTTP handlers
- No SQL inside application layer
- No framework dependency inside domain layer
- Proper error handling
- Context propagation
- Structured DTO usage

---

## Technical Requirements

- Use Fiber for HTTP
- Use sqlx for database access
- Use PostgreSQL 15+
- Use UUID for primary keys
- Use environment variables for config
- Provide:
  - Dockerfile
  - docker-compose.yml
- App must run with:

  docker compose up --build

- App must auto-run migrations at startup

---

## Expected Output

AI must generate:

- Full project structure
- All Go files
- Migration SQL file
- Dockerfile
- docker-compose.yml
- README with run instructions

Application must compile and run successfully.
