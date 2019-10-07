package domain

import "time"

type ErrorMsg struct {
	ProcessId    string    `json:"processId"`
	SensorUuid   string    `json:"sensorUuid"`
	Service      string    `json:"service"`
	ErrorMsg     string    `json:"errorMsg"`
	ErrorDetails string    `json:"errorDetails"`
	PublishedOn  time.Time `json:"publishedOn"`
}
