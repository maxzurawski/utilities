package domain

import "time"

type LogMsg struct {
	ProcessId   string    `json:"processId"`
	SensorUuid  string    `json:"sensorUuid"`
	Service     string    `json:"service"`
	LogMsg      string    `json:"logMsg"`
	PublishedOn time.Time `json:"publishedOn"`
}
