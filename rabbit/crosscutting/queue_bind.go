package crosscutting

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func (r *RabbitConnector) BindQueue(queue amqp.Queue, routingKey, exchangeName string) {

	err := r.Channel.QueueBind(queue.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		log.Panic(fmt.Sprintf("could not bind queue: [%s] with routingkey: [%s] on exchangeName: [%s]", queue.Name, routingKey, exchangeName))
	}
}
