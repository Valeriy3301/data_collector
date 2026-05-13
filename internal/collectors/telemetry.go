package collectors

import (
	"errors"
	"math/rand"
	"time"

	"your-module/internal/infra"
	"your-module/internal/models"
)

type TelemetryCollector struct {
	rl    *infra.RateLimiter
	retry *infra.RetryHandler
}

func NewTelemetryCollector(rl *infra.RateLimiter, retry *infra.RetryHandler) *TelemetryCollector {
	return &TelemetryCollector{rl: rl, retry: retry}
}

func (t *TelemetryCollector) Fetch() ([]models.DataPoint, error) {
	if !t.rl.Allow() {
		return nil, errors.New("rate limited")
	}

	var out []models.DataPoint

	err := t.retry.Do(func() error {
		if rand.Intn(10) < 3 {
			return errors.New("telemetry api error")
		}

		out = []models.DataPoint{
			{
				Source: "telemetry",
				Type:   "CPU",
				Value:  rand.Float64() * 100,
				Time:   time.Now().Unix(),
			},
		}
		return nil
	})

	return out, err
}