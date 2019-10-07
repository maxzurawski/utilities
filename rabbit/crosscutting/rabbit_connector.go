package crosscutting

import "github.com/streadway/amqp"

type RabbitConnector struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func InitConnector(connection *amqp.Connection) *RabbitConnector {
	rabbitConnector := &RabbitConnector{
		Conn: connection,
	}
	rabbitConnector.initChannel()
	return rabbitConnector
}

func (r *RabbitConnector) Reset() {
	_ = r.Channel.Close()
	r.initChannel()
	r.DeclareQueue()
}
