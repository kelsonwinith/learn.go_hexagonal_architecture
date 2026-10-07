# Repository Guidelines

This project is a Go learning project for hexagonal architecture. Keep changes aligned with the existing dependency direction:

`domain` defines business models, errors, and ports. `application` implements use cases and depends on domain ports. `adapter/in` handles external input such as Fiber HTTP. `adapter/out` handles external services such as PostgreSQL. `bootstrap` and module `Init` functions wire concrete dependencies together.

## Project Structure

```text
.
├── bin/                                  # compiled application binaries
├── cmd/
│   └── main.go                           # application entry point
├── docs/                                 # generated Swagger docs
└── internal/
    ├── bootstrap/                        # startup, infra setup (Config, DB), middleware, module init
    │   └── app.go
    ├── infrastructure/                   # config loading, DB connection, migrations, seed, persistence models
    │   ├── config/
    │   └── postgresql/
    ├── modules/
    │   ├── exampleBasic/                 # baseline scaffold: flat CRUD with one output port
    │   │   ├── domain/                   # business model, domain errors, input/output ports
    │   │   ├── application/              # use cases; depend on domain ports, not concrete adapters
    │   │   ├── adapter/
    │   │   │   ├── in/
    │   │   │   │   └── fiber/            # HTTP handlers and request/response DTOs
    │   │   │   └── out/
    │   │   │       └── postgresql/       # GORM adapters and domain/model mappers
    │   │   └── exampleBasic_module.go    # wires the module's dependencies
    │   └── exampleAdvanced/              # advanced reference: relations, transaction, multiple output ports
    │       ├── domain/
    │       ├── application/
    │       ├── adapter/
    │       │   ├── in/
    │       │   │   └── fiber/
    │       │   └── out/
    │       │       ├── postgresql/
    │       │       └── eventlog/         # second output adapter, not a database
    │       └── exampleAdvanced_module.go
    └── shared/                           # reusable domain errors and shared adapter helpers
        ├── domain/
        └── adapter/
            ├── in/
            │   └── fiber/
            └── out/
                └── postgresql/
```

## Architecture Rules

- Preserve dependency flow from outside to inside:
  - Fiber handlers call use case ports.
  - Use cases call domain-defined output ports.
  - PostgreSQL adapters implement domain-defined output ports.
  - Domain code must not import Fiber, GORM, PostgreSQL, config, or infrastructure packages.
- Put new business rules in domain constructors/methods or application use cases, not HTTP handlers or database adapters.
- Keep DTOs and persistence models out of the domain layer. Convert at adapter boundaries with DTO `ToDomain` helpers or mapper functions.
- Register new dependencies explicitly in the module initializer, following `internal/modules/exampleBasic/exampleBasic_module.go` (baseline) or `internal/modules/exampleAdvanced/exampleAdvanced_module.go` (multiple output ports).
- Use `context.Context` through use case and adapter methods, matching the existing `Execute(ctx, ...)` pattern.

## Example Module Showcase

The `example*` modules are the reference implementations. Pick the one whose shape matches the feature you are adding, then use its matching use case as a template.

### Example Basic (`internal/modules/exampleBasic`)

Flat, single-table CRUD with one output port. This is the baseline scaffold to copy when starting a new module.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleUsecaseCreate` | `POST /api/v1/examplebasic` | Simple write: validation through the `NewExample` domain constructor, then a single insert |
| `ExampleUsecaseCreateMultiple` | `POST /api/v1/examplebasic/batch` | Atomic batch insert wrapped in `ExamplePostgresqlTransaction.WithinTransaction` |
| `ExampleUsecaseGetAll` | `GET /api/v1/examplebasic` | Simple read delegated straight to the output port |
| `ExampleUsecaseGetPaginated` | `GET /api/v1/examplebasic/paginated` | Pagination + search: query DTO validation, shared `sharedDomain.NewPagination` default/max rule, then `sharedFiber.ResponsePaginated` |
| `ExampleUsecaseGetByID` | `GET /api/v1/examplebasic/:id` | Read with `gorm.ErrRecordNotFound` mapped to `ExampleErrNotFound` in the adapter |
| `ExampleUsecaseUpdate` | `PUT /api/v1/examplebasic/:id` | Ownership rule (`ExampleErrForbidden`) plus domain mutation via `UpdateExample` |
| `ExampleUsecaseDelete` | `DELETE /api/v1/examplebasic/:id` | Ownership rule plus GORM soft delete (`deleted_by` then `Delete`) |

### Example Advanced (`internal/modules/exampleAdvanced`)

Aggregate relations and more than one output port. Copy from here when a feature spans multiple tables or must talk to a non-database dependency.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleAdvancedUsecaseCreate` | `POST /api/v1/exampleadvanced` | Parent + children created atomically inside `WithinTransaction`, then a domain event published through a second output port (`eventlog`) |
| `ExampleAdvancedUsecaseGetByID` | `GET /api/v1/exampleadvanced/:id` | Loading an aggregate with children via GORM `Preload`, with not-found mapped to `ExampleAdvancedErrNotFound` |

## Naming Conventions

- Follow the existing file naming style: `<module>_<layer>_<name>.go`. The name is exactly three underscore-separated segments, and each segment may use camelCase, for example `exampleBasic_usecase_create.go`, `exampleAdvanced_postgresql_getByID.go`, `model_postgresql_exampleAdvancedParent.go`, and `exampleAdvanced_domain_model.go`.
- Apply the same `<module>_<layer>_<name>` convention to folders, which may also use camelCase, for example `internal/modules/exampleAdvanced/adapter/in/fiber/exampleAdvanced_fiber_getByID.go`.
- Constructors should be named `New<Type>` and return the interface when exposing a port implementation from the application layer.
- Use `Execute` for use case and output adapter methods, as defined by the domain port interfaces.
- Keep import aliases consistent with the project style, such as `exampleBasicDomain`, `exampleAdvancedPostgresql`, and `sharedFiber`.

## Imports

- Always alias every import with an explicit name, even when it matches the package name, so each reference is unambiguous:

  ```go
  import (
  	errors "errors"
  	reflect "reflect"
  	strings "strings"

  	validator "github.com/go-playground/validator/v10"
  	fiber "github.com/gofiber/fiber/v3"
  	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
  )
  ```

- Group imports into three blocks separated by blank lines, in this order: standard library, third-party, then internal (`github.com/kelsonwinith/...`) packages.
- Use lowercase aliases for standard library and third-party packages (`fiber`, `gorm`, `validator`).
- Use camelCase aliases for internal packages, prefixed with the module or layer name (`exampleBasicDomain`, `sharedFiber`, `exampleAdvancedPostgresql`); a bare package name such as `config` is fine when unambiguous.
- Never rely on the implicit package name; write the alias on every import.

## File Sectioning

Every Go file is divided into labeled banner sections so declarations are easy to locate. Each banner uses this exact three-line format, where both rule lines are the fixed string `// ============================================================================` and the title is replaced with the section name:

```go
// ============================================================================
// <Title>
// ============================================================================
```

Place sections in the order below, omitting any section that would be empty:

1. `Constants` - `const` blocks.
2. `Variables` - package-level `var` blocks.
3. `Types` - structs, interfaces, type aliases, and enums.
4. `Constructors` - `New<Type>` functions.
5. `Methods` - functions with a receiver.
6. `Functions` - package-level helpers.

Rules:

- Separate a banner from the declarations that follow it with one blank line, and precede it with one blank line when it is not the first item after the imports.
- A file that contains only one category still gets that category's banner (a port-only file starts with a `Types` banner).
- When a category has meaningful subgroups, give each subgroup its own named banner in the same format, such as `Usecase Ports` and `PostgreSQL Ports` in `exampleBasic_domain_port.go`.
- Keep the Swagger doc comment directly above its handler: place the `Methods` banner, then a blank line, then the `// Handle ...` annotation block, so the annotations stay attached to `func`.
- Do not section declarations inside a function body. Existing inline step comments (for example `// Adapters Out - PostgreSQL` in `exampleBasic_module.go`) stay as they are.

Example:

```go
package domain

import (
	context "context"
)

// ============================================================================
// Constants
// ============================================================================

const ExampleMaxLength = 255

// ============================================================================
// Types
// ============================================================================

type Example struct {
	ID string
}

// ============================================================================
// Constructors
// ============================================================================

func NewExample(id string) *Example {
	return &Example{ID: id}
}

// ============================================================================
// Methods
// ============================================================================

func (e *Example) Validate() error {
	return nil
}

// ============================================================================
// Functions
// ============================================================================

func helper() string {
	return ""
}
```

## HTTP and Validation

- Fiber handlers should:
  - Bind input with `sharedFiber.Bind[URI, Query, Body](c)`. Use `sharedFiber.Empty` for parts not required.
  - Convert DTOs to domain models before calling use cases.
  - Return responses with `sharedFiber.ResponseSuccess`, `ResponseCreated`, `ResponseNoContent`, or `ResponseError`.
- Add validation tags to request DTOs where needed. The shared validator is configured in `internal/bootstrap/app.go` and supports `json`, `query`, `params`, and `uri` tags.
- Reusable response envelopes (for example pagination via `sharedFiber.ResponsePaginated`) belong in `internal/shared/adapter/in/fiber`, not in a module DTO package. Map the domain page items to response DTOs first, then pass them with the `sharedDomain.Pagination` and total.
- Use `sharedDomain.NewPagination` for page/page-size defaulting and max clamping instead of reimplementing it in each use case. Pass `sharedDomain.PaginationLimits` to override the shared defaults; omit it to fall back to `DefaultPageSize`/`MaxPageSize`.
- Keep Swagger comments on handlers up to date when adding or changing endpoints.

## Error Handling

- Use `sharedDomain.Error` for application-level errors.
- Errors consist of an `HTTPCode`, a `Type` (e.g., `BAD_REQUEST`, `NOT_FOUND`), a unique `ID` (e.g., `E001`), and a `Message`.
- Use prefix-based IDs defined in `internal/shared/domain/shared_domain_error.go` (e.g., `SYS` for system, `FIB` for fiber, `EXB` for the exampleBasic module, `EXA` for the exampleAdvanced module).
- Prefer defining reusable errors in the domain layer of the relevant module.

## PostgreSQL and GORM

- Use GORM as the ORM. The database is configured with `SingularTable: true`.
- Use the shared PostgreSQL wrapper from `internal/shared/adapter/out/postgresql`.
- Use `GetExecutor(ctx)` so transactional contexts continue to work.
- Convert between domain objects and GORM models in `adapter/out/postgresql/mapper`.
- Keep persistence models to one file per database table (e.g. `model_postgresql_exampleAdvancedParent.go` and `model_postgresql_exampleAdvancedChild.go`).
- For multi-step writes that must be atomic, follow the existing transaction pattern used by `ExampleUsecaseCreateMultiple` (single table) or `ExampleAdvancedUsecaseCreate` (parent + children across multiple output ports).

## Configuration

- Environment variables are defined in `internal/infrastructure/config/config_schema.go` via struct tags (`envconfig`, `default`, `required`).
- Whenever you add, rename, or remove a config field in `config_schema.go`, update `.env.example` in the same change so it lists the matching variable name.
- `.env.example` declares variable names only (empty values), never real values. `.env` is local and git-ignored; never commit credentials.

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
- For HTTP binding/response behavior, use `httptest` with `app.Test`.
- Prefer testing use cases with small fake port implementations instead of a real database unless persistence behavior is the subject of the test.
- Run `make test` before handing off code changes.

## Code Quality

- Keep changed Go files formatted. If formatting requires a command, use a `Makefile` target for it; otherwise report that formatting could not be run under the command rule.
- Keep changes scoped to the relevant module/layer.
- Do not introduce new frameworks or infrastructure abstractions unless the existing structure cannot support the requested behavior.
- When adding environment-driven config, update `config_schema.go` and `.env.example` together (see Configuration).
