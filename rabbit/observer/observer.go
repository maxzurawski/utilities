package observer

import (
	"github.com/streadway/amqp"
	"github.com/xdevices/utilities/rabbit/crosscutting"
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
