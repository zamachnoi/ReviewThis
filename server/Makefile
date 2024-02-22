restart: build
	@docker compose up -d

build:
	@docker build -t viewthis-backend .

start:
	@docker-compose up -d

stop:
	@docker-compose down

go:
	@go build -o bin/viewthis-backend cmd/api/main.go