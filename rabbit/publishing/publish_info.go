package publishing

import (
	"time"

	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/domain"
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
