package main

import (
	"log"
	"time"

	"concurrent-collector/internal/config"
	"concurrent-collector/internal/collectors"
	"concurrent-collector/internal/concurrency"
	"concurrent-collector/internal/infra"
)

func main() {
	log.Println("Starting Concurrent Data Collector...")

	cfg := config.DefaultConfig()

	rl := infra.NewRateLimiter(cfg.RequestsPerSecond)
	retry := infra.NewRetryHandler(cfg.MaxRetries, cfg.RetryDelay)

	crypto := collectors.NewCryptoCollector(rl, retry)
	telemetry := collectors.NewTelemetryCollector(rl, retry)

	batcher := concurrency.NewBatcher(cfg.BatchSize, cfg.BatchInterval)
	pool := concurrency.NewWorkerPool(cfg.Workers, batcher)

	pool.Start()

	go runCollector(crypto, pool, 2*time.Second)
	go runCollector(telemetry, pool, 3*time.Second)

	select {}
}

func runCollector(c collectors.Collector, pool *concurrency.WorkerPool, interval time.Duration) {
	for {
		data, err := c.Fetch()
		if err == nil {
			pool.Submit(data)
		}
		time.Sleep(interval)
	}
}