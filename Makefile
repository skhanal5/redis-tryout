.PHONY: vet fmt build test

vet:
	go vet

fmt:
	go fmt ./...

build: clean
	go build cmd/redis-tryout/main.go

test:
	go test ./...
 