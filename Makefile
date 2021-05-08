.PHONY: build
build:
	rm -rf build && go build -o build/go-cron-parser .

format:
	go clean && go fmt ./... && go vet ./...

test:
	go test -v ./...