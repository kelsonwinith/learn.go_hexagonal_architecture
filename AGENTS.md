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
    │   ├── exampleUser/                  # baseline scaffold: single-entity CRUD, auth and ownership rules
    │   │   ├── domain/                   # business model, domain errors, input/output ports
    │   │   ├── application/              # use cases; depend on domain ports, not concrete adapters
    │   │   ├── adapter/
    │   │   │   ├── in/
    │   │   │   │   └── fiber/            # HTTP handlers and request/response DTOs
    │   │   │   └── out/
    │   │   │       └── postgresql/       # GORM adapters and domain/model mappers
    │   │   └── exampleUser_module.go     # wires the module's dependencies
    │   ├── exampleProduct/               # single-entity CRUD with a numeric (price) field
    │   │   ├── domain/
    │   │   ├── application/
    │   │   ├── adapter/
    │   │   │   ├── in/
    │   │   │   │   └── fiber/
    │   │   │   └── out/
    │   │   │       └── postgresql/
    │   │   └── exampleProduct_module.go
    │   └── exampleOrder/                 # aggregate, transaction, multiple output ports, cross-module services
    │       ├── domain/
    │       ├── application/
    │       ├── adapter/
    │       │   ├── in/
    │       │   │   └── fiber/
    │       │   └── out/
    │       │       ├── postgresql/
    │       │       ├── eventlog/         # second output adapter, not a database
    │       │       ├── exampleuser/      # cross-module adapter consuming the exampleUser service
    │       │       └── exampleproduct/   # cross-module adapter consuming the exampleProduct service
    │       └── exampleOrder_module.go
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
- Register new dependencies explicitly in the module initializer, following `internal/modules/exampleUser/exampleUser_module.go` (baseline) or `internal/modules/exampleOrder/exampleOrder_module.go` (multiple output ports and cross-module services).
- Cross-module calls go through a port defined in the consumer's `domain`; the provider exposes its use case and a small `adapter/out/<provider>` in the consumer maps the result. Wire them in `bootstrap`, e.g. `exampleOrder` consuming the `exampleUser` and `exampleProduct` services.
- Use `context.Context` through use case and adapter methods, matching the existing `Execute(ctx, ...)` pattern.

## Example Module Showcase

The `example*` modules are the reference implementations. Pick the one whose shape matches the feature you are adding, then use its matching use case as a template.

### Example User (`internal/modules/exampleUser`)

Flat, single-table CRUD with one output port. This is the baseline scaffold to copy when starting a new module.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleUserUsecaseCreate` | `POST /api/v1/exampleuser` | Simple write: validation through the `NewExampleUser` domain constructor, then a single insert |
| `ExampleUserUsecaseCreateMultiple` | `POST /api/v1/exampleuser/batch` | Atomic batch insert wrapped in `ExampleUserPostgresqlTransaction.WithinTransaction` |
| `ExampleUserUsecaseGetAll` | `GET /api/v1/exampleuser` | Simple read delegated straight to the output port |
| `ExampleUserUsecaseGetPaginated` | `GET /api/v1/exampleuser/paginated` | Pagination + search: query DTO validation, shared `sharedDomain.NewPagination` default/max rule, then `sharedFiber.ResponsePaginated` |
| `ExampleUserUsecaseGetByID` | `GET /api/v1/exampleuser/:id` | Read with `gorm.ErrRecordNotFound` mapped to `ExampleUserErrNotFound` in the adapter |
| `ExampleUserUsecaseUpdate` | `PUT /api/v1/exampleuser/:id` | Ownership rule (`ExampleUserErrForbidden`) plus domain mutation via `UpdateExampleUser` |
| `ExampleUserUsecaseDelete` | `DELETE /api/v1/exampleuser/:id` | Ownership rule plus GORM soft delete (`deleted_by` then `Delete`) |

### Example Product (`internal/modules/exampleProduct`)

Same single-entity CRUD shape as `exampleUser`, adding a numeric `Price` field with its own domain validation. Copy from here when an entity has numeric or money-like attributes.

### Example Order (`internal/modules/exampleOrder`)

Aggregate relations, a transaction, more than one output port, and cross-module services. Copy from here when a feature spans multiple tables, talks to a non-database dependency, or consumes another module.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleOrderUsecaseCreate` | `POST /api/v1/exampleorder` | Order + products created atomically inside `WithinTransaction`; the referenced user and products are resolved through the `ExampleOrderUserReader` and `ExampleOrderProductReader` ports (cross-module), then a domain event is published through the `eventlog` output port |
| `ExampleOrderUsecaseGetByID` | `GET /api/v1/exampleorder/:id` | Loading an aggregate with products via GORM `Preload`, with not-found mapped to `ExampleOrderErrNotFound` |

## Naming Conventions

- Follow the existing file naming style: `<module>_<layer>_<name>.go`. The name is exactly three underscore-separated segments, and each segment may use camelCase, for example `exampleUser_usecase_create.go`, `exampleOrder_postgresql_getByID.go`, `model_postgresql_exampleOrder.go`, and `exampleOrder_domain_model.go`.
- Apply the same `<module>_<layer>_<name>` convention to folders, which may also use camelCase, for example `internal/modules/exampleOrder/adapter/in/fiber/exampleOrder_fiber_getByID.go`.
- Constructors should be named `New<Type>` and return the interface when exposing a port implementation from the application layer.
- Use `Execute` for use case and output adapter methods, as defined by the domain port interfaces.
- Keep import aliases consistent with the project style, such as `exampleUserDomain`, `exampleOrderPostgresql`, and `sharedFiber`.
- Prefix every identifier and variable declared inside a module with the module name: exported names use PascalCase (`ExampleUserUsecaseCreate`, `ExampleOrderPostgresqlCreate`) and unexported or local names use camelCase (`exampleUserPostgresqlCreate`, `exampleOrderCreatePostgres`). Shared packages are the exception and use the `shared` prefix instead, for example `sharedDomain`, `sharedFiber`, and `sharedPostgresql`.

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
- Use camelCase aliases for internal packages, prefixed with the module or layer name (`exampleUserDomain`, `sharedFiber`, `exampleOrderPostgresql`); a bare package name such as `config` is fine when unambiguous.
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
- When a category has meaningful subgroups, give each subgroup its own named banner in the same format, such as `Usecase Ports` and `PostgreSQL Ports` in `exampleUser_domain_port.go`.
- Keep the Swagger doc comment directly above its handler: place the `Methods` banner, then a blank line, then the `// Handle ...` annotation block, so the annotations stay attached to `func`.
- Do not section declarations inside a function body. Existing inline step comments (for example `// Adapters Out - PostgreSQL` in `exampleUser_module.go`) stay as they are.

Example:

```go
package domain

import (
	context "context"
)

// ============================================================================
// Constants
// ============================================================================

const ExampleUserMaxLength = 255

// ============================================================================
// Types
// ============================================================================

type ExampleUser struct {
	ID string
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUser(id string) *ExampleUser {
	return &ExampleUser{ID: id}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUser) Validate() error {
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
- Use prefix-based IDs defined in `internal/shared/domain/shared_domain_error.go` (e.g., `SYS` for system, `FIB` for fiber, `EXU` for the exampleUser module, `EXP` for the exampleProduct module, `EXO` for the exampleOrder module).
- Prefer defining reusable errors in the domain layer of the relevant module.

## PostgreSQL and GORM

- Use GORM as the ORM. The database is configured with `SingularTable: true`.
- Use the shared PostgreSQL wrapper from `internal/shared/adapter/out/postgresql`.
- Use `GetExecutor(ctx)` so transactional contexts continue to work.
- Convert between domain objects and GORM models in `adapter/out/postgresql/mapper`.
- Keep persistence models to one file per database table (e.g. `model_postgresql_exampleOrder.go` and `model_postgresql_exampleOrderProduct.go`).
- For multi-step writes that must be atomic, follow the existing transaction pattern used by `ExampleUserUsecaseCreateMultiple` (single table) or `ExampleOrderUsecaseCreate` (order + products across multiple output ports).

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
