package domain

import "time"

type TemperatureMeasurement struct {
	ProcessId   string    `json:"processId"`
	SensorUuid  string    `json:"sensorUuid"`
	Service     string    `json:"service"`
	Value       float64   `json:"value"`
	PublishedOn time.Time `json:"publishedOn"`
}
