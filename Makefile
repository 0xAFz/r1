BINARY_NAME=r1
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GORUN=$(GO) run
GOFMT=$(GO) fmt
GOMOD=$(GO) mod

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "-w -s" .

run:
	$(GORUN) main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

fmt:
	$(GOFMT) ./...

deps:
	$(GOMOD) download

install:
	$(GOMOD) tidy

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

lint:
	golint ./...

vet:
	$(GO) vet ./...

check: lint vet test

dev:
	$(GORUN) main.go

start: build
	./$(BINARY_NAME)

.PHONY: all build run test clean fmt deps install coverage lint vet check dev start
