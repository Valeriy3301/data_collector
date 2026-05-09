package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Starting Concurrent Data Collector...")

	cfg := DefaultConfig()

	rateLimiter := NewRateLimiter(cfg.RequestsPerSecond)
	retry := NewRetryHandler(cfg.MaxRetries, cfg.RetryDelay)

	cryptoCollector := NewCryptoCollector(rateLimiter, retry)
	telemetryCollector := NewTelemetryCollector(rateLimiter, retry)

	batcher := NewBatcher(cfg.BatchSize, cfg.BatchInterval)
	workerPool := NewWorkerPool(cfg.Workers, batcher)

	workerPool.Start()

	// Crypto pipeline
	go func() {
		for {
			data, err := cryptoCollector.Fetch()
			if err == nil {
				workerPool.Submit(data)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	// Telemetry pipeline
	go func() {
		for {
			data, err := telemetryCollector.Fetch()
			if err == nil {
				workerPool.Submit(data)
			}
			time.Sleep(3 * time.Second)
		}
	}()

	select {}
}