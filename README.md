# Concurrent Data Collector (Go)

A high-performance concurrent data ingestion service built in Go

## Features

- Goroutine-based concurrency
- Worker pool architecture
- Rate limiting (token bucket)
- Retry mechanism with backoff
- Batch processing
- Mock data collectors (crypto + telemetry)
- Docker support

## Run

```bash
go run main.go
```

---

## Docker

```bash
docker build -t collector .
docker run collector
```

## TODO

- Prometheus metrics
- OpenTelemetry tracing
- Kafka ingestion pipeline
- graceful shutdown (signals) (!!!)