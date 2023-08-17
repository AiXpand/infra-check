# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Build parameters
BINARY_NAME = infra-checker
LINUX_BINARY = $(BINARY_NAME)_linux_amd64
MACOS_BINARY = $(BINARY_NAME)_macos_arm64

# Main build target
all: clean build

build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./dist/$(LINUX_BINARY) ./cmd
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ./dist/$(MACOS_BINARY) ./cmd

clean:
	$(GOCLEAN)
	rm -f ./dist/$(LINUX_BINARY)
	rm -f ./dist/$(MACOS_BINARY)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd
	./$(BINARY_NAME)

deps:
	$(GOGET) ./...

.PHONY: all build clean test run deps
