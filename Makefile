APP_NAME = dapr-work

build:
	@echo "==> Building..."
	cd pub-app && go build -o pub-app main.go
	cd sub-app && go build -o sub-app main.go

run: build
	@echo "==> Running..."
	cd pub-app && ./pub-app &
	cd sub-app && ./sub-app

dapr-run:
	@echo "==> Running Dapr..."
	(cd pub-app && start cmd /c "dapr run --app-id pub-app --app-port 8080 --dapr-http-port 3500 --resources-path ../.dapr/components --config ../.dapr/config.yaml --app-protocol http --log-level fatal -- go run ./main.go")
	(cd sub-app && start cmd /c "dapr run --app-id sub-app --app-port 8081 --dapr-http-port 3501 --resources-path ../.dapr/components --config ../.dapr/config.yaml --app-protocol http --log-level fatal -- go run ./main.go")

dapr-stop:
	@echo "==> Stopping Dapr..."
	dapr stop --app-id pub-app
	dapr stop --app-id sub-app