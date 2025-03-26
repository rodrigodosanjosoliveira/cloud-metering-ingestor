.PHONY: build run test swag

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
