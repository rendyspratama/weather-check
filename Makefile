# Project Variables
BINARY_NAME = weather-app
BUILD_DIR = ./build_dir
GO_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Go Commands
.PHONY: all build run clean tidy fmt lint help

all: build

## Build the Go binary
build:
	@echo "Building the project..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete! Binary is at $(BUILD_DIR)/$(BINARY_NAME)"

## Run the application
run:
	@echo "Running the project..."
	go run main.go

## Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Cleanup complete!"

## Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed!"

## Format and Lint
fmt:
	@echo "Formatting Go files..."
	go fmt $(GO_FILES)

lint:
	@echo "Running Go vet..."
	go vet $(GO_FILES)

## Help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build       Build the Go binary"
	@echo "  run         Run the application"
	@echo "  clean       Remove build artifacts"
	@echo "  deps        Install and tidy dependencies"
	@echo "  fmt         Format Go files"
	@echo "  lint        Run Go vet"
