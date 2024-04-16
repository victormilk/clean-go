run:
	@echo "Running the application"
	@go run cmd/app/main.go

build:
	@echo "Building the application"
	@go build -o bin/app cmd/app/main.go cmd/app/wire_gen.go

run-build: build
	@echo "Running the application"
	@./bin/app

coverage:
	@echo "Running the tests with coverage"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

MIGRATE_NAME ?= "initial"

create-migrate:
	@echo "Applying the migrations"
	@atlas migrate diff $(MIGRATE_NAME) \
       --dir "file://migrations" \
       --to "file://schema.sql" \
       --dev-url "docker://postgres?search_path=public"

.PHONY: run build run-build coverage