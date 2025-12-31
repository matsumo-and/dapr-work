.PHONY: proto run-auth run-user build-auth build-user tidy docker-build-auth docker-build-user test-isolation

proto:
	buf generate

run-auth:
	go run cmd/auth-service/main.go

run-user:
	go run cmd/user-service/main.go

build-auth:
	go build -o bin/auth-service cmd/auth-service/main.go

build-user:
	go build -o bin/user-service cmd/user-service/main.go

build-all: build-auth build-user

tidy:
	go mod tidy

docker-build-auth:
	docker build -t auth-service:latest -f Dockerfile.auth .

docker-build-user:
	docker build -t user-service:latest -f Dockerfile.user .

docker-build-all: docker-build-auth docker-build-user

# サービス間の依存関係がないことをテスト
test-isolation:
	@./scripts/check-isolation.sh
