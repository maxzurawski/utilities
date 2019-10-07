package crosscutting

import "fmt"

func (r *RabbitConnector) DeclareTopicExchange(exchangeName string) {
	err := r.Channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(fmt.Sprintf("could not declare exchange %s", exchangeName))
	}
}
