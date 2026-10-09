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
    ├── bootstrap/                        # startup wiring, one file per topic
    │   ├── bootstrap_config_init.go      # config loading
    │   ├── bootstrap_database_init.go    # DB connection, migrations, seed, close
    │   ├── bootstrap_app_init.go         # Fiber app + validator
    │   ├── bootstrap_middleware_init.go  # logger, CORS, swagger
    │   ├── bootstrap_module_init.go      # module wiring incl. cross-module adapters
    │   └── bootstrap_server_run.go       # listen + graceful shutdown
    ├── infrastructure/                   # config loading, DB connection, migrations, seed, persistence models
    │   ├── config/
    │   └── postgresql/
    ├── modules/
    │   ├── exampleUser/                  # identity module: register, login, profile
    │   │   ├── domain/                   # business model, domain errors, input/output ports
    │   │   ├── application/              # use cases; depend on domain ports, not concrete adapters
    │   │   ├── adapter/
    │   │   │   ├── in/
    │   │   │   │   └── fiber/            # HTTP handlers and request/response DTOs
    │   │   │   └── out/
    │   │   │       └── postgresql/       # GORM adapters and domain/model mappers
    │   │   └── exampleUser_module.go     # wires the module's dependencies
    │   ├── exampleProduct/               # baseline scaffold: single-entity CRUD with one output port
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
    │       │       ├── exampleUser/      # cross-module adapter consuming the exampleUser service
    │       │       └── exampleProduct/   # cross-module adapter consuming the exampleProduct service
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
- Register new dependencies explicitly in the module initializer, following `internal/modules/exampleProduct/exampleProduct_module.go` (baseline) or `internal/modules/exampleOrder/exampleOrder_module.go` (multiple output ports and cross-module services).
- Cross-module calls go through a port defined in the consumer's `domain`; the provider exposes its use case and a small `adapter/out/<provider>` in the consumer maps the result. Wire them in `bootstrap`, e.g. `exampleOrder` consuming the `exampleUser` and `exampleProduct` services.
- Use `context.Context` through use case and adapter methods, matching the existing `Execute(ctx, ...)` pattern.

## Example Module Showcase

The `example*` modules are the reference implementations. Pick the one whose shape matches the feature you are adding, then use its matching use case as a template.

### Example User (`internal/modules/exampleUser`)

Identity module (register, login, profile). Keep it small: it exists to show auth flow and to be consumed by other modules, not to show CRUD.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleUserUsecaseRegister` | `POST /api/v1/exampleuser/register` | Validation through the `NewExampleUser` domain constructor (name, email, password), then a single insert |
| `ExampleUserUsecaseLogin` | `POST /api/v1/exampleuser/login` | Login: lookup by email through `ExampleUserPostgresqlGetByEmail`, bcrypt password check (`ExampleUserPasswordComparer`), then a signed JWT issued through `sharedDomain.TokenService` |
| `ExampleUserUsecaseGetByID` | `GET /api/v1/exampleuser/:id` | Read with `gorm.ErrRecordNotFound` mapped to `ExampleUserErrNotFound`; also the service `exampleOrder` consumes cross-module |

### Example Product (`internal/modules/exampleProduct`)

Single-entity CRUD with one output port. Copy from here when starting a new module with create, update, delete, get-by-ID and paginated read.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleProductUsecaseCreate` | `POST /api/v1/exampleproduct` | Simple write: validation (incl. numeric `Price`) through `NewExampleProduct`, then a single insert |
| `ExampleProductUsecaseGetByID` | `GET /api/v1/exampleproduct/:id` | Read with `gorm.ErrRecordNotFound` mapped to `ExampleProductErrNotFound` |
| `ExampleProductUsecaseGetPaginated` | `GET /api/v1/exampleproduct/paginated` | Pagination + search: query DTO validation, shared `sharedDomain.NewPagination` default/max rule, then `sharedFiber.ResponsePaginated` |
| `ExampleProductUsecaseUpdate` | `PUT /api/v1/exampleproduct/:id` | Ownership rule (`ExampleProductErrForbidden`) plus domain mutation via `UpdateExampleProduct` |
| `ExampleProductUsecaseDelete` | `DELETE /api/v1/exampleproduct/:id` | Ownership rule plus GORM soft delete (`deleted_by` then `Delete`) |

### Example Order (`internal/modules/exampleOrder`)

Aggregate relations, a transaction, more than one output port, and cross-module services. Copy from here when a feature spans multiple tables, talks to a non-database dependency, or consumes another module.

| Use case (`application/`) | Endpoint | Showcases |
| --- | --- | --- |
| `ExampleOrderUsecaseCreate` | `POST /api/v1/exampleorder` | Order + products created atomically inside `WithinTransaction`; the referenced user and products are resolved through the `ExampleOrderUserReader` and `ExampleOrderProductReader` ports (cross-module), then a domain event is published through the `eventlog` output port |
| `ExampleOrderUsecaseGetByID` | `GET /api/v1/exampleorder/:id` | Loading an aggregate with products via GORM `Preload`, with not-found mapped to `ExampleOrderErrNotFound` |
| `ExampleOrderUsecaseGetPaginated` | `GET /api/v1/exampleorder/paginated` | Paginated aggregate listing: shared `sharedDomain.NewPagination` + `sharedFiber.ResponsePaginated` |

## Naming Conventions

- Follow the existing file naming style: `<module>_<layer>_<name>.go`. The name is exactly three underscore-separated segments, and each segment may use camelCase, for example `exampleUser_usecase_register.go`, `exampleOrder_postgresql_getByID.go`, `model_postgresql_exampleOrder.go`, and `exampleOrder_domain_model.go`.
- Apply the same `<module>_<layer>_<name>` convention to folders, which may also use camelCase, for example `internal/modules/exampleOrder/adapter/in/fiber/exampleOrder_fiber_getByID.go`.
- Constructors should be named `New<Type>` and return the interface when exposing a port implementation from the application layer.
- Use `Execute` for use case and output adapter methods, as defined by the domain port interfaces.
- Keep import aliases consistent with the project style, such as `exampleUserDomain`, `exampleOrderPostgresql`, and `sharedFiber`.
- Prefix every identifier declared inside a module with the module token: exported names use PascalCase (`ExampleUserUsecaseRegister`, `ExampleOrderPostgresqlCreate`) and unexported names use camelCase.

### Variable Naming

- Prefix every struct field, constructor parameter, and local variable with the module token: `exampleProductCreatePostgres`, `exampleOrderUserReader`.
- Name injected collaborators as `<module><Operation><Adapter>`: `exampleProductCreatePostgres`, `exampleOrderEventPublisher`, `exampleUserPasswordHasher`, `exampleOrderUserReader`.
- Name domain values `<module>` (singular), `<module>s` (slice), or `<module><Sub>` (sub-entity): `exampleOrder`, `exampleOrders`, `exampleOrderProduct`, `exampleProductInput`.
- Keep Go idioms and primitive/audit parameters unprefixed: `ctx`, `err`, `ok`, `i`, `id`, `email`, `password`, `createdBy`, `updatedBy`, `page`, `pageSize`, `search`, `total`, `pagination`.
- Keep receivers short (`uc`, `h`, `e`, `p`, `r`); name DTO method receivers `request`.
- Shared packages are the exception and use the `shared` prefix for their import aliases, for example `sharedDomain`, `sharedFiber`, and `sharedPostgresql`.

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
- Add validation tags to request DTOs where needed. The shared validator is configured in `internal/bootstrap/bootstrap_app_init.go` and supports `json`, `query`, `params`, and `uri` tags.
- Reusable response envelopes (for example pagination via `sharedFiber.ResponsePaginated`) belong in `internal/shared/adapter/in/fiber`, not in a module DTO package. Map the domain page items to response DTOs first, then pass them with the `sharedDomain.Pagination` and total.
- Use `sharedDomain.NewPagination` for page/page-size defaulting and max clamping instead of reimplementing it in each use case. Pass `sharedDomain.PaginationLimits` to override the shared defaults; omit it to fall back to `DefaultPageSize`/`MaxPageSize`.
- Authentication is JWT bearer. `exampleUser.Login` issues an HS256 token through `sharedDomain.TokenService` (implemented by `internal/shared/adapter/out/jwt`), and `sharedFiber.NewAuth` validates the `Authorization: Bearer <token>` header. Read the authenticated subject with `sharedFiber.GetAuthUserID(c)` (the user UUID string) and use it as the domain `CreatedBy`/`UpdatedBy` and for ownership checks; never read auth headers directly in handlers.
- Keep Swagger comments on handlers up to date when adding or changing endpoints.

### Swagger Annotation Template

Place this block directly above every `Handle` method: the `Methods` banner, a blank line, then the annotation comment so it stays attached to the function. Copy it and adjust names; keep the tag, DTO references, and router in the forms shown.

```go
// Handle CreateExampleProduct
// @Summary Create an example product
// @Description Create a new example product
// @Tags Example Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param example body exampleProductDto.ExampleProductCreateRequest true "Create ExampleProduct"
// @Success 201 {object} exampleProductDto.ExampleProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct [post]
func (h *ExampleProductFiberCreate) Handle(c fiber.Ctx) error {
```

Rules:

- `@Tags` is the only thing that controls the Swagger section header. Write it in Title Case with spaces (`Example User`, `Example Product`, `Example Order`), never the lowercase module name, and use the exact same value on every handler of a module. Swag splits tags on commas, not spaces, so a spaced name stays a single tag.
- `@Router` keeps the lowercase module path (`/api/v1/exampleproduct`) even though the tag is `Example Product`.
- Use the module DTO package alias for types: `exampleProductDto.ExampleProductResponse`.
- For paginated reads use `@Success 200 {object} sharedFiber.ResponsePaginatedData[exampleProductDto.ExampleProductResponse]`.
- Add `@Security BearerAuth` only on endpoints guarded by the auth middleware. Do not add a per-route `Authorization` header `@Param`: the Swagger UI Authorize button (from the single `@securityDefinitions.apikey BearerAuth` block in `cmd/main.go`) supplies `Authorization: Bearer <token>` for every secured operation.
- A module-wide default is not used on purpose: a global `@security BearerAuth` in `cmd/main.go` would also mark public endpoints (register, login, reads) as secured.

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
- For multi-step writes that must be atomic, follow the transaction pattern used by `ExampleOrderUsecaseCreate` (order + products across multiple output ports).

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
