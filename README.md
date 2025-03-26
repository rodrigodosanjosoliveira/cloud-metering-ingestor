[![Go CI](https://github.com/rodrigodosanjosoliveira/cloud-metering-ingestor/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/rodrigodosanjosoliveira/cloud-metering-ingestor/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/rodrigodosanjosoliveira/cloud-metering-ingestor/graph/badge.svg?token=DOFQTSDUL3)](https://codecov.io/gh/rodrigodosanjosoliveira/cloud-metering-ingestor)

# Cloud Metering Ingestor

Microservice for ingesting and aggregating usage pulses for a cloud billing platform. Built with Go, Clean Architecture, validation, Swagger documentation, containerization, and CI.

---

## Features

- Pulse ingestion via REST API
- In-memory aggregation by tenant, SKU, and unit
- Simulated publish on flush using a `Publisher` interface
- Structured logging with Zap
- Validation using `validator.v10`
- Unit testing with Testify
- Swagger documentation
- Dockerized with Docker Compose
- CI pipeline with GitHub Actions
- Linting with golangci-lint

---

## Technologies

- Go 1.23
- Gin Web Framework
- Uber Zap Logger
- Go Validator
- Swaggo for Swagger generation
- Docker + Docker Compose
- GitHub Actions (CI)
- golangci-lint

---

## API Endpoints

### `POST /pulses`
Receive a new usage pulse.

```json
{
  "tenant": "tenant_xpto",
  "product_sku": "sku_vm_123",
  "used_amount": 307,
  "use_unit": "KB",
  "timestamp": "2025-03-21T10:30:00Z"
}
```

### `GET /aggregates`
Returns the current usage aggregation by tenant + SKU + unit.

### `POST /flush`
Simulates publishing aggregated data and clears the internal state.

---

## Documentation

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Swagger spec: [`docs/swagger.yaml`](docs/swagger.yaml)
- Architecture Diagram: [`docs/diagram.puml`](docs/diagram.puml)

---

## Running with Docker

```bash
docker compose up --build
```

Then access: [http://localhost:8080](http://localhost:8080)

### Running Load Test
```bash
docker run --rm -i --network cloud-metering-ingestor_default grafana/k6 run - < load-test/pulse_test.js
```

---

## Code Quality

- Run linter:

```bash
golangci-lint run
```

- Run tests:

```bash
go test ./...
```

---

## Setup locally (manual)

```bash
go mod tidy
go run cmd/main.go
```

---

## About the Challenge

This project was developed as part of a technical challenge for a backend position. It focuses on high-quality Go code, microservice architecture principles, observability, and development best practices.

**Assumptions made:**
- Pulses are aggregated in-memory for simplicity
- Flush simulates integration with a downstream service using a mock publisher
- Authentication/authorization was not in scope

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
