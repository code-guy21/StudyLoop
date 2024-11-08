.PHONY: dev test build deploy

dev:
	docker-compose up --build

test:
	cd server && go test ./...
	cd client && npm test

build:
	docker-compose build

deploy:
	echo "Deployment configuration to be added"

.PHONY: migrate

migrate:
	cd server && go run cmd/migrate/main.go