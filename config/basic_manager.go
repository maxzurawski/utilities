package config

import (
	"github.com/maxzurawski/utilities/discovery"
)

type BasicManager interface {
	DBPath() string
	ServiceName() string
	RegistrationTicket() *discovery.RegistrationTicket
	EurekaService() string
	Address() string
}
