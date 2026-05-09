package main

import (
	"golang.org/x/time/rate"
	"time"
)

type RateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter(rps int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Every(time.Second/time.Duration(rps)), rps),
	}
}

func (r *RateLimiter) Allow() bool {
	return r.limiter.Allow()
}