.PHONY: proto run build docker-build

proto:
	buf generate

run:
	go run cmd/main.go

build:
	go build -o bin/user-service cmd/main.go

docker-build:
	docker build -t user-service:latest .

tidy:
	go mod tidy
