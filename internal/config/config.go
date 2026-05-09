package main

import "time"

type Config struct {
	Workers           int
	BatchSize         int
	BatchInterval     time.Duration
	RequestsPerSecond int
	MaxRetries        int
	RetryDelay        time.Duration
}

func DefaultConfig() Config {
	return Config{
		Workers:           5,
		BatchSize:         10,
		BatchInterval:     5 * time.Second,
		RequestsPerSecond: 10,
		MaxRetries:        3,
		RetryDelay:        500 * time.Millisecond,
	}
}