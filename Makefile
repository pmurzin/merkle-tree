.PHONY: all build run test clean

all: build

build:
	go build -o bin/merkle-tree cmd/main.go

run: build
	./bin/merkle-tree

test:
	go test  -v ./internal/merkle

clean:
	rm -rf bin
