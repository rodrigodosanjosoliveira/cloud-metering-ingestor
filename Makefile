.PHONY: deps build run test swag diagram coverage

deps:
	go mod tidy
	go mod vendor

build:
	go build -o bin/ingestor cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

swag:
	swag init -g cmd/main.go -o docs

diagram:
	docker run --rm -v $(PWD)/docs:/workspace plantuml/plantuml diagram.puml

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out
	firefox coverage.html