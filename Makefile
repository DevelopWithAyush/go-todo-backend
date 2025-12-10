# Go Todo App Makefile
# Industrial standard commands for development and deployment

.PHONY: all build run dev clean test swagger deps help

# Default target
all: swagger build

# Build the application
build:
	@echo "Building application..."
	go build -o ./tmp/main.exe ./cmd/api

# Run the application
run: build
	@echo "Running application..."
	./tmp/main.exe

# Run with hot reload using air
dev:
	@echo "Starting development server with hot reload..."
	air

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf ./tmp/main.exe
	rm -rf ./docs

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Generate swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
	@echo "Swagger docs generated at /swagger/index.html"

# Install/update swagger CLI tool
swagger-install:
	@echo "Installing swag CLI..."
	go install github.com/swaggo/swag/cmd/swag@latest

# Format swagger documentation (regenerate with latest changes)
swagger-fmt:
	@echo "Formatting Swagger documentation..."
	swag fmt -g cmd/api/main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Update dependencies
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Help
help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Build and run the application"
	@echo "  make dev            - Run with hot reload (air)"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make test           - Run tests"
	@echo "  make swagger        - Generate Swagger documentation"
	@echo "  make swagger-install - Install swag CLI tool"
	@echo "  make swagger-fmt    - Format swagger annotations"
	@echo "  make deps           - Install dependencies"
	@echo "  make deps-update    - Update dependencies"
	@echo "  make help           - Show this help message"
