package publishing

import (
	"time"

	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/domain"
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
