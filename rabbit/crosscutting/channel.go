package crosscutting

func (r *RabbitConnector) initChannel() {
	channel, err := r.Conn.Channel()
	if err != nil {
		panic("cannot create channel")
	}
	r.Channel = channel
}
