.PHONY: run build test docker-up docker-down migrate-up migrate-down

BINARY_NAME=tutorlink

run:
	@go run ./cmd/api

build:
	@go build -o bin/${BINARY_NAME} ./cmd/api

test:
	@go test -v ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

migrate-up:
	@migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/tutorlink?sslomode=disable" up

migrate-down:
	@migrate -path migrations -database "postgres://posttgres:postgres@localhost:5432/tutorlink?sslmode=disable" down

deps:
	@go mod download
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest