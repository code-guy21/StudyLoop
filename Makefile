.PHONY: dev test build deploy

dev:
	docker-compose up --build

test:
	cd backend && go test ./...
	cd frontend && npm test

build:
	docker-compose build

deploy:
	echo "Deployment configuration to be added"

.PHONY migrate

migrate:
	cd backend && go run cmd/migrate/main.go