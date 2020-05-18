# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=go_roller
BINARY_UNIX=$(BINARY_NAME)_unix
TAG=go_roller:latest
D=docker

.PHONY: build test clean run roll

all: test build
build:
	@$(GOBUILD) -o $(BINARY_NAME) -v
test:
	@$(GOTEST) -v ./...
clean:
	@rm -rf $(BINARY_NAME)
run:
	@$(GOBUILD) -o $(BINARY_NAME) -v
	@./$(BINARY_NAME)
deploy:
	@$(D) build -t $(TAG) .
	@$(D) run -p 8080:8080 --rm $(TAG)
