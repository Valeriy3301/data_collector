package models

type DataPoint struct {
	Source string
	Type   string
	Value  float64
	Time   int64
}