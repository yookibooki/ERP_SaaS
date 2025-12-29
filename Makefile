.PHONY: all fmt build clean

all: fmt build

fmt:
	gofumpt -l -w .
	goimports -l -w .
	go mod tidy

build:
	mkdir -p bin
	go build -o bin/_binary_name_ ./cmd/_your_app_

clean:
	rm -rf bin
