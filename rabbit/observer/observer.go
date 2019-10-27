package observer

import (
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/streadway/amqp"
)

type Observer struct {
	*crosscutting.RabbitConnector
	Queue amqp.Queue
}

func InitObserver(connection *amqp.Connection) *Observer {
	observer := &Observer{
		RabbitConnector: crosscutting.InitConnector(connection),
	}
	observer.Queue = observer.DeclareQueue()

	return observer
}

func (o *Observer) ShutdownObserver() {
	_ = o.Channel.Close()
}
