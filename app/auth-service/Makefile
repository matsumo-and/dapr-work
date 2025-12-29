.PHONY: proto run build docker-build

proto:
	buf generate

run:
	go run cmd/main.go

build:
	go build -o bin/auth-service cmd/main.go

docker-build:
	docker build -t auth-service:latest .

tidy:
	go mod tidy
