run: build
	@./bin/reviewthis-backend

build:
	@go build -o bin/reviewthis-backend cmd/api/main.go

up:
	@go run cmd migrate/main.go