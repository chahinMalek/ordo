.PHONY: all build test clean install fmt vet

BINARY_NAME=ordo
DIST_DIR=dist

all: build

build:
	@mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(BINARY_NAME) ./cmd/ordo

test:
	go test -v ./...

clean:
	rm -rf $(DIST_DIR)
	go clean

install:
	go install ./cmd/ordo

fmt:
	go fmt ./...

vet:
	go vet ./...
