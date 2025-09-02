# note: call scripts from /scripts

.PHONY: run build tidy

run:
	@echo "Starting inventory-api with live reload..."
	@air

build:
	@echo "Building inventory-api binary..."
	@go build -o ./build/inventory-api ./cmd/inventory-api

tidy:
	@go mod tidy