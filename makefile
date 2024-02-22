build:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

run:
	go run cmd/server/main.go

generate:
	go run cmd/migrate/generate.go

migrate:
	go run cmd/migrate/migrate.go

rollback:
	go run cmd/migrate/rollback.go

format:
	go fmt ./...