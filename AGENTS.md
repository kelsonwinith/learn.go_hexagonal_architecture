# Repository Guidelines

This project is a Go learning project for hexagonal architecture. Keep changes aligned with the existing dependency direction:

`domain` defines business models, errors, and ports. `application` implements use cases and depends on domain ports. `adapter/in` handles external input such as Fiber HTTP. `adapter/out` handles external services such as PostgreSQL. `bootstrap` and module `Init` functions wire concrete dependencies together.

## Project Structure

- `bin/`: compiled application binaries.
- `cmd/main.go`: application entry point.
- `internal/bootstrap`: application startup, infrastructure setup (Config, DB), middleware, and module initialization.
- `internal/modules/example/domain`: domain model, domain errors, and input/output port interfaces.
- `internal/modules/example/application`: use case implementations. Use cases should depend on domain interfaces, not concrete adapters.
- `internal/modules/example/adapter/in/fiber`: HTTP handlers and request/response DTOs for Fiber.
- `internal/modules/example/adapter/out/postgresql`: PostgreSQL adapter implementations using GORM and domain/model mappers.
- `internal/infrastructure`: config loading, database connection setup, migrations, seed data, and persistence models.
- `internal/shared`: reusable domain errors, shared adapter helpers (Fiber, PostgreSQL).

## Architecture Rules

- Preserve dependency flow from outside to inside:
  - Fiber handlers call use case ports.
  - Use cases call domain-defined output ports.
  - PostgreSQL adapters implement domain-defined output ports.
  - Domain code must not import Fiber, GORM, PostgreSQL, config, or infrastructure packages.
- Put new business rules in domain constructors/methods or application use cases, not HTTP handlers or database adapters.
- Keep DTOs and persistence models out of the domain layer. Convert at adapter boundaries with DTO `ToDomain` helpers or mapper functions.
- Register new dependencies explicitly in the module initializer, following `internal/modules/example/example_module.go`.
- Use `context.Context` through use case and adapter methods, matching the existing `Execute(ctx, ...)` pattern.

## Naming Conventions

- Follow the existing file naming style: `<module>_<layer>_<operation>.go`, for example `example_usecase_create.go` and `example_postgresql_getByID.go`.
- Constructors should be named `New<Type>` and return the interface when exposing a port implementation from the application layer.
- Use `Execute` for use case and output adapter methods, as defined by the domain port interfaces.
- Keep import aliases consistent with the project style, such as `exampleDomain`, `sharedFiber`, and `examplePostgresql`.

## HTTP and Validation

- Fiber handlers should:
  - Bind input with `sharedFiber.Bind[URI, Query, Body](c)`. Use `sharedFiber.Empty` for parts not required.
  - Convert DTOs to domain models before calling use cases.
  - Return responses with `sharedFiber.ResponseSuccess`, `ResponseCreated`, `ResponseNoContent`, or `ResponseError`.
- Add validation tags to request DTOs where needed. The shared validator is configured in `internal/bootstrap/app.go` and supports `json`, `query`, `params`, and `uri` tags.
- Keep Swagger comments on handlers up to date when adding or changing endpoints.

## Error Handling

- Use `sharedDomain.Error` for application-level errors.
- Errors consist of an `HTTPCode`, a `Type` (e.g., `BAD_REQUEST`, `NOT_FOUND`), a unique `ID` (e.g., `E001`), and a `Message`.
- Use prefix-based IDs defined in `internal/shared/domain/shared_domain_error.go` (e.g., `SYSM` for system, `FIB` for fiber, `E` for example module).
- Prefer defining reusable errors in the domain layer of the relevant module.

## PostgreSQL and GORM

- Use GORM as the ORM. The database is configured with `SingularTable: true`.
- Use the shared PostgreSQL wrapper from `internal/shared/adapter/out/postgresql`.
- Use `GetExecutor(ctx)` so transactional contexts continue to work.
- Convert between domain objects and GORM models in `adapter/out/postgresql/mapper`.
- For multi-step writes that must be atomic, follow the existing transaction pattern used by `ExampleUsecaseCreateMultiple`.

## Commands

- Only run commands through `make` targets defined in `Makefile`. Do not run raw `go`, `docker-compose`, Swagger, or other project commands directly unless the user explicitly asks for one.

- Run all tests:
  ```sh
  make test
  ```

- Generate Swagger docs:
  ```sh
  make swagger
  ```

- Start only PostgreSQL:
  ```sh
  make db-up
  ```

- Start all services (App + DB):
  ```sh
  make compose-up
  ```

- Stop all services:
  ```sh
  make compose-down
  ```

- Build the application:
  ```sh
  make build
  ```

- Run the application locally (generates Swagger and starts DB):
  ```sh
  make run
  ```

## Testing Guidance

- Add focused tests near the package being changed.
- For HTTP binding/response behavior, follow the style in `internal/shared/adapter/in/fiber/*_test.go` with `httptest` and `app.Test`.
- Prefer testing use cases with small fake port implementations instead of a real database unless persistence behavior is the subject of the test.
- Run `make test` before handing off code changes.

## Code Quality

- Keep changed Go files formatted. If formatting requires a command, use a `Makefile` target for it; otherwise report that formatting could not be run under the command rule.
- Keep changes scoped to the relevant module/layer.
- Do not introduce new frameworks or infrastructure abstractions unless the existing structure cannot support the requested behavior.
- When adding environment-driven config, update `internal/infrastructure/config` and any required local setup documentation together.
