.PHONY: vet fmt build test

vet:
	go vet

fmt:
	go fmt ./...

build:
	go build -o redis-tryout cmd/redis-tryout/main.go

test:
	go test ./...
 