package domain

import "time"

type CacheUpdateStatus struct {
	ProcessId   string    `json:"processId"`
	Service     string    `json:"service"`
	ErrorMsg    string    `json:"errorMsg"`
	Cache       string    `json:"cache"`
	PublishedOn time.Time `json:"publishedOn"`
}
