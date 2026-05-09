package main

import "log"

type WorkerPool struct {
	workers int
	input   chan []DataPoint
	batcher *Batcher
}

func NewWorkerPool(workers int, batcher *Batcher) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		input:   make(chan []DataPoint, 100),
		batcher: batcher,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.workers; i++ {
		go w.worker(i)
	}
}

func (w *WorkerPool) Submit(data []DataPoint) {
	w.input <- data
}

func (w *WorkerPool) worker(id int) {
	for data := range w.input {
		log.Printf("Worker %d processing %d items\n", id, len(data))

		for _, d := range data {
			w.batcher.Add(d)
		}
	}
}