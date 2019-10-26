package rabbit

import (
	"github.com/maxzurawski/utilities/rabbit/observer"
	"github.com/maxzurawski/utilities/rabbit/publishing"
	"github.com/streadway/amqp"
)

type RabbitMQManager struct {
	conn *amqp.Connection
}

func (r *RabbitMQManager) InitConnection(url string) {
	connection, err := amqp.Dial(url)
	if err != nil {
		panic("cannot connect to rabbitmq")
	}
	r.conn = connection
}

func (r *RabbitMQManager) CloseConnection() {
	r.conn.Close()
}

func (r *RabbitMQManager) InitPublisher() *publishing.Publisher {
	return publishing.InitPublisher(r.conn)
}

func (r *RabbitMQManager) InitObserver() *observer.Observer {
	return observer.InitObserver(r.conn)
}
