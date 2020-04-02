# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=go_roller
BINARY_UNIX=$(BINARY_NAME)_unix

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
roll:
	@./$(BINARY_NAME) --dicetoroll 20 --amountofdice 1
	@./$(BINARY_NAME) --dicetoroll 20 --amountofdice 1 --modifier 6
	@./$(BINARY_NAME) --dicetoroll 20 --amountofdice 2 --modifier 8 --advantage
	@./$(BINARY_NAME) --dicetoroll 20 --amountofdice 2 --modifier -2 --disadvantage 
	@./$(BINARY_NAME) --dicetoroll 6 --amountofdice 6 --modifier 3
