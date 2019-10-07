package crosscutting

import (
	"github.com/streadway/amqp"

	"github.com/labstack/gommon/log"
)

func (r *RabbitConnector) DeclareQueue() amqp.Queue {
	queue, err := r.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Panic("could not declare queue")
	}

	return queue
}
