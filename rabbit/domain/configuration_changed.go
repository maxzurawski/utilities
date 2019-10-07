package domain

import "time"

type ConfigurationChanged struct {
	ProcessId   string    `json:"processId"`
	Service     string    `json:"service"`
	Previous    string    `json:"previous"`
	Current     string    `json:"current"`
	PublishedOn time.Time `json:"publishedOn"`
}
