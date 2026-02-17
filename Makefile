.PHONY: compose-up compose-down down swagger run migration-up migration-down db-up

# Docker Compose Up
compose-up:
	docker-compose up --build --remove-orphans

# Docker Compose Down
compose-down:
	docker-compose down

# Run Migrations Up
migration-up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/hexagonal_go?sslmode=disable" -verbose up

# Run Migrations Down
migration-down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/hexagonal_go?sslmode=disable" -verbose down

# Generate Swagger Docs
swagger:
	go run github.com/swaggo/swag/cmd/swag init -g cmd/main.go --parseDependency --parseInternal

# Start Database only
db-up:
	docker-compose up -d db --remove-orphans

# Run Application (Generates Swagger docs first, ensures DB is up)
run: swagger db-up
	go run cmd/main.go
