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
	return &CryptoCollector{rl: rl, retry: retry}
}

func (c *CryptoCollector) Fetch() ([]models.DataPoint, error) {
	if !c.rl.Allow() {
		return nil, errors.New("rate limited")
	}

	var out []models.DataPoint

	err := c.retry.Do(func() error {
		if rand.Intn(10) < 2 {
			return errors.New("crypto API error")
		}

		out = []models.DataPoint{
			{
				Source: "crypto",
				Type:   "BTC",
				Value:  rand.Float64() * 60000,
				Time:   time.Now().Unix(),
			},
		}
		return nil
	})

	return out, err
}