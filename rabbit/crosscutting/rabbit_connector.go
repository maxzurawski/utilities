package crosscutting

import "github.com/streadway/amqp"

type RabbitConnector struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
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
	if r.Queue == nil {
		queue := r.DeclareQueue()
		r.Queue = &queue
	}

}
