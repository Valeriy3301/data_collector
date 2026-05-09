package collectors

import "concurrent-collector/internal/models"

type Collector interface {
	Fetch() ([]models.DataPoint, error)
}