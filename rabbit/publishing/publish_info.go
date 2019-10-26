package publishing

import (
	"time"

	"../crosscutting"
	"../domain"
)

func (d *Publisher) PublishInfo(
	processId,
	sensorUuid,
	service,
	logMsg string) {

	now := time.Now()

	msg := &domain.LogMsg{
		ProcessId:   processId,
		SensorUuid:  sensorUuid,
		Service:     service,
		LogMsg:      logMsg,
		PublishedOn: now,
	}

	d.publishLogMsg(*msg, crosscutting.RoutingKeyLogsSeverityInfo.String())
}
