package publishing

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/domain"
)

func (d *Publisher) PublishTemperatureMeasurement(
	routingKey,
	service,
	sensorUuid string,
	value float64) {

	now := time.Now()

	msg := domain.TemperatureMeasurement{
		ProcessId:   uuid.New().String(),
		Value:       value,
		Service:     service,
		PublishedOn: now,
		SensorUuid:  sensorUuid,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Warn("could not marshal msg")
		return
	}
	d.Reset()
	d.Publish(
		crosscutting.TopicMeasurements.String(),
		routingKey,
		string(bytes))

	d.Reset()
	d.PublishExtInfo(msg)
}
