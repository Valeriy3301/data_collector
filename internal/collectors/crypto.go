package collectors

import (
	"errors"
	"math/rand"
	"time"

	"concurrent-collector/internal/infra"
	"concurrent-collector/internal/models"
)

type CryptoCollector struct {
	rl    *infra.RateLimiter
	retry *infra.RetryHandler
}

func NewCryptoCollector(rl *infra.RateLimiter, retry *infra.RetryHandler) *CryptoCollector {
	return &CryptoCollector{
		rl:    rl,
		retry: retry,
	}
}

func (c *CryptoCollector) Fetch() ([]models.DataPoint, error) {
	// 1. rate limiting
	if !c.rl.Allow() {
		return nil, errors.New("rate limited")
	}

	var out []models.DataPoint

	// 2. retry wrapper
	err := c.retry.Do(func() error {

		// имитация нестабильного API
		if rand.Intn(10) < 2 {
			return errors.New("crypto API error")
		}

		// mock data (в будущем заменишь на real HTTP call)
		out = []models.DataPoint{
			{
				Source: "crypto",
				Type:   "BTC",
				Value:  50000 + rand.Float64()*10000,
				Time:   time.Now().Unix(),
			},
			{
				Source: "crypto",
				Type:   "ETH",
				Value:  3000 + rand.Float64()*500,
				Time:   time.Now().Unix(),
			},
		}

		return nil
	})

	return out, err
}