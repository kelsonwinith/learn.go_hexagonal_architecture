# Docker Compose Up
compose-up:
	docker-compose up --build --remove-orphans

# Docker Compose Down
compose-down:
	docker-compose down

# Generate Swagger Docs
swagger:
	$$(go env GOPATH)/bin/swag init -g cmd/main.go --parseDependency --parseInternal

# Start Database only
db-up:
	docker-compose up -d db --remove-orphans

# Test
test:
	go test -v ./...

# Build Application
build:
	go build -o bin/app cmd/main.go

# Run Application (Generates Swagger docs first, ensures DB is up)
run: swagger db-up
	go run cmd/main.go
