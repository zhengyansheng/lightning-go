run:
	@go run cmd/server/main.go --config=config/config.toml

run-prod:
	@go run cmd/server/main.go --config=config/config-prod.toml

migrate:
	@go run cmd/migrate/main.go --config=config/config.toml

migrate-prod:
	@go run cmd/migrate/main.go --config=config/config-prod.toml

swagger:
	@swag init -g cmd/server/main.go

build-server-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/lightning-go cmd/server/main.go

build-windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/lightning-go cmd/server/main.go
