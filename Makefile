APP_NAME = dapr-work

build:
	@echo "==> Building..."
	cd pub-app && go build -o pub-app main.go
	cd sub-app && go build -o sub-app main.go

run: build
	@echo "==> Running..."
	cd pub-app && ./pub-app &
	cd sub-app && ./sub-app

docker-compose-up:
	@echo "==> Starting Docker Compose..."
	docker compose up -d

docker-compose-down:
	@echo "==> Stopping Docker Compose..."
	docker compose down

dapr-init:
	@echo "==> Initializing Dapr..."
	dapr uninstall
	dapr init --slim

dapr-run: dapr-init docker-compose-up
	@echo "==> Running Dapr..."
	(cd pub-app && dapr run --app-id pub-app --app-port 8080 --dapr-http-port 3500 --resources-path ../.dapr/components --config ../.dapr/config.yaml --app-protocol http --log-level fatal -- go run ./main.go) &
	(cd sub-app && dapr run --app-id sub-app --app-port 8081 --dapr-http-port 3501 --resources-path ../.dapr/components --config ../.dapr/config.yaml --app-protocol http --log-level fatal -- go run ./main.go) &

dapr-stop: docker-compose-down
	@echo "==> Stopping Dapr..."
	dapr stop --app-id pub-app
	dapr stop --app-id sub-app