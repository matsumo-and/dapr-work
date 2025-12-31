.PHONY: proto build-auth build-user build-all tidy docker-build-auth docker-build-user docker-build-all \
	docker-up docker-down docker-logs docker-logs-db docker-clean \
	dapr-init dapr-run-auth dapr-run-user \
	db redis test-isolation help

#######################
# Proto & Go Build
#######################

proto:
	buf generate

build-auth:
	go build -o bin/auth-service cmd/auth-service/main.go

build-user:
	go build -o bin/user-service cmd/user-service/main.go

build-all: build-auth build-user

tidy:
	go mod tidy

#######################
# Docker Build
#######################

docker-build-auth:
	docker build -t auth-service:latest -f Dockerfile.auth .

docker-build-user:
	docker build -t user-service:latest -f Dockerfile.user .

docker-build-all: docker-build-auth docker-build-user

#######################
# Infrastructure (Docker Compose)
#######################

docker-up:
	docker-compose up -d
	@echo "✅ Infrastructure started!"
	@echo "- PostgreSQL: localhost:5432"
	@echo ""
	@echo "Next steps:"
	@echo "  1. Initialize Dapr (first time only): make dapr-init"
	@echo "     (Redis, Zipkin, Placement service are auto-installed)"
	@echo "  2. Run services with Dapr:"
	@echo "     make dapr-run-auth"
	@echo "     make dapr-run-user"

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-logs-db:
	docker-compose logs -f postgres

docker-clean:
	docker-compose down -v

#######################
# Dapr
#######################

dapr-init:
	dapr init
	@echo "✅ Dapr initialized!"

dapr-run-auth:
	dapr run \
		--app-id auth-service \
		--app-port 8080 \
		--dapr-http-port 3500 \
		--dapr-grpc-port 50001 \
		--components-path ./.dapr/components \
		--log-level info \
		-- go run cmd/auth-service/main.go

dapr-run-user:
	dapr run \
		--app-id user-service \
		--app-port 8081 \
		--dapr-http-port 3501 \
		--dapr-grpc-port 50002 \
		--components-path ./.dapr/components \
		--log-level info \
		-- go run cmd/user-service/main.go

#######################
# Database Access
#######################

db:
	docker exec -it dapr-postgres psql -U app -d appdb

redis:
	docker exec -it dapr_redis redis-cli

#######################
# Testing
#######################

test-isolation:
	@./scripts/check-isolation.sh

#######################
# Help
#######################

help:
	@echo "Available commands:"
	@echo ""
	@echo "Proto & Build:"
	@echo "  make proto              - Generate proto code with buf"
	@echo "  make build-auth         - Build auth-service binary"
	@echo "  make build-user         - Build user-service binary"
	@echo "  make build-all          - Build all service binaries"
	@echo "  make tidy               - Run go mod tidy"
	@echo ""
	@echo "Docker Build:"
	@echo "  make docker-build-auth  - Build auth-service Docker image"
	@echo "  make docker-build-user  - Build user-service Docker image"
	@echo "  make docker-build-all   - Build all Docker images"
	@echo ""
	@echo "Infrastructure:"
	@echo "  make docker-up          - Start infrastructure (PostgreSQL)"
	@echo "  make docker-down        - Stop infrastructure"
	@echo "  make docker-logs        - Show logs for all infrastructure"
	@echo "  make docker-logs-db     - Show PostgreSQL logs"
	@echo "  make docker-clean       - Remove all containers and volumes"
	@echo ""
	@echo "Dapr:"
	@echo "  make dapr-init          - Initialize Dapr (Redis, Zipkin, Placement auto-installed)"
	@echo "  make dapr-run-auth      - Run auth-service with Dapr"
	@echo "  make dapr-run-user      - Run user-service with Dapr"
	@echo ""
	@echo "Database:"
	@echo "  make db                 - Connect to PostgreSQL"
	@echo "  make redis              - Connect to Dapr Redis"
	@echo ""
	@echo "Testing:"
	@echo "  make test-isolation     - Test service isolation"
	@echo ""
	@echo "Help:"
	@echo "  make help               - Show this help message"
