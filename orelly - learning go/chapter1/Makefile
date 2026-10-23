.DEFAULT_GOAL := all
.PHONY: fmt vet build run clean all

all: run clean

fmt:
	@echo "==> Formatting..."
	go fmt ./...

vet: fmt
	@echo "==> Checking errors..."
	go vet ./...

build: vet
	@echo "==> Building..."
	go build

run: build
	@echo "==> Running..."
	./hello-world

clean:
	@echo "==> Cleaning..."
	rm -f hello-world
