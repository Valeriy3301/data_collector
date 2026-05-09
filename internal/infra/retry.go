package main

import "time"

type RetryHandler struct {
	maxRetries int
	delay      time.Duration
}

func NewRetryHandler(max int, delay time.Duration) *RetryHandler {
	return &RetryHandler{
		maxRetries: max,
		delay:      delay,
	}
}

func (r *RetryHandler) Do(fn func() error) error {
	var err error

	for i := 0; i < r.maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(r.delay)
	}

	return err
}