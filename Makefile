
build:
	go build -o bin/main ./cmd/server/main.go

run: build
	./bin/main

test:
	go test -v ./...

