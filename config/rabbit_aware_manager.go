package config

import "github.com/maxzurawski/utilities/rabbit/publishing"

type RabbitAwareManager interface {
	BasicManager
	ConnectToRabbit() bool
	InitPublisher() *publishing.Publisher
}
