# buld the app
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# run the app
run:
	@go run cmd/api/main.go

# test
test:
	@echo "Testing..."
	@go test ./... -v 

# clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: all build run test clean watch
