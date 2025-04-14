.PHONY: vet fmt build test

vet:
	go vet

fmt:
	go fmt ./...

build: fmt
	go build -o redis-tryout cmd/redis-tryout/main.go

test: fmt
	go test ./...
 