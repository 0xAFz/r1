# Variables
BINARY_NAME=kumo
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GORUN=$(GO) run
GOFMT=$(GO) fmt
GOMOD=$(GO) mod

# Default target
all: build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "-w -s" .

# Run the project
run:
	$(GORUN) main.go

# Test the project
test:
	$(GOTEST) -v ./...

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Format the code
fmt:
	$(GOFMT) ./...

# Download dependencies
deps:
	$(GOMOD) download

# Install dependencies
install:
	$(GOMOD) tidy

# Run all tests and generate coverage report
coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# Run the linter
lint:
	golint ./...

# Run the vet tool
vet:
	$(GO) vet ./...

# Run all checks (lint, vet, test)
check: lint vet test

# Run the application in development mode
dev:
	$(GORUN) main.go

# Build and run the application
start: build
	./$(BINARY_NAME)

# Phony targets
.PHONY: all build run test clean fmt deps install coverage lint vet check dev start
