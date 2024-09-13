.PHONY: gen-env lint test build clean

gen-env:
	cp ./config/.env.default .env

lint:
	go mod tidy
	golangci-lint run ./... --timeout=10m

test:
	go test ./... -cover

build:
	go mod tidy
	go build -o build/main cmd/main.go

clean:
	rm -rf build