package publishing

import (
	"github.com/streadway/amqp"
	"github.com/xdevices/utilities/rabbit/crosscutting"
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
