package domain

import (
	"time"

	"github.com/xdevices/utilities/stringutils"
)

type NotifierMsg struct {
	ProcessId   string                  `json:"processId"`
	SensorUuid  string                  `json:"sensorUuid"`
	Service     string                  `json:"service"`
	Max         string                  `json:"max"`
	Min         string                  `json:"min"`
	Value       stringutils.MultiString `json:"value"`
	LogMsg      string                  `json:"logMsg"`
	PublishedOn time.Time               `json:"publishedOn"`
}
