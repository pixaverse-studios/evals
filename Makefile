.PHONY: all build test clean lint

# Variables
BINARY_NAME=eval-cli
BUILD_DIR=bin
MAIN_PATH=cmd/eval-cli/main.go

all: clean build test

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

test:
	@echo "Running tests..."
	@go test ./... -v

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean

lint:
	@echo "Linting..."
	@golangci-lint run

install:
	@echo "Installing..."
	@go install $(MAIN_PATH)

coverage:
	@echo "Generating coverage report..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

deps:
	@echo "Downloading dependencies..."
	@go mod download

update-deps:
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy