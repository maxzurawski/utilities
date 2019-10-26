package publishing

import (
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/streadway/amqp"
)

type Publisher struct {
	*crosscutting.RabbitConnector
}

func InitPublisher(connection *amqp.Connection) *Publisher {
	dialer := &Publisher{
		RabbitConnector: crosscutting.InitConnector(connection),
	}
	return dialer
}
