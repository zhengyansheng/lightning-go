run:
	@go run cmd/server/main.go --config=config/server.toml

migrate:
	@go run cmd/migrate/main.go --config=config/migrate.toml

migrate-prod:
	@go run cmd/migrate/main.go --config=config/migrate-prod.toml

swagger:
	@swag init -g cmd/server/main.go

build-migrate-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/lightning-migrate-go cmd/migrate/main.go

build-server-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/lightning-go cmd/server/main.go

build-windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/lightning-go cmd/server/main.go
