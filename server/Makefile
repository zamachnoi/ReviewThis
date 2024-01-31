run: build
	@./bin/viewthis-backend

build:
	@go build -o bin/viewthis-backend cmd/api/main.go

up:
	@go run cmd migrate/main.go