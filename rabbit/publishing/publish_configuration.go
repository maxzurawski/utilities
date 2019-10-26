package publishing

import (
	"encoding/json"
	"time"

	"../crosscutting"
	"../domain"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func (d *Publisher) PublishConfigurationChanged(
	routingKey,
	service string,
	previousValue,
	currentValue interface{}) {

	previous, _ := json.Marshal(previousValue)
	current, _ := json.Marshal(currentValue)

	now := time.Now()

	msg := domain.ConfigurationChanged{
		ProcessId:   uuid.New().String(),
		Service:     service,
		Previous:    string(previous),
		Current:     string(current),
		PublishedOn: now,
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Warn("could not marshal msg")
		return
	}
	d.Publish(
		crosscutting.TopicConfigurationChanged.String(),
		routingKey,
		string(bytes))

	d.PublishExtInfo(msg)
}
