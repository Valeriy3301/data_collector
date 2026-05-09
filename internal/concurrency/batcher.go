package main

import (
	"log"
	"sync"
	"time"
)

type Batcher struct {
	size     int
	interval time.Duration

	mu    sync.Mutex
	batch []DataPoint
}

func NewBatcher(size int, interval time.Duration) *Batcher {
	b := &Batcher{
		size:     size,
		interval: interval,
		batch:    make([]DataPoint, 0, size),
	}

	go b.flushLoop()
	return b
}

func (b *Batcher) Add(dp DataPoint) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.batch = append(b.batch, dp)

	if len(b.batch) >= b.size {
		b.flush()
	}
}

func (b *Batcher) flushLoop() {
	ticker := time.NewTicker(b.interval)

	for range ticker.C {
		b.mu.Lock()
		b.flush()
		b.mu.Unlock()
	}
}

func (b *Batcher) flush() {
	if len(b.batch) == 0 {
		return
	}

	log.Printf("Flushing batch size=%d\n", len(b.batch))

	// simulate persistence / external API
	b.batch = make([]DataPoint, 0, b.size)
}