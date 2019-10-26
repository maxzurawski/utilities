package publishing

import (
	"time"

	"github.com/google/uuid"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/domain"
)

func (d *Publisher) PublishError(
	processId,
	sensorUuid,
	service,
	errorMsg, errorDetails string) {

	now := time.Now()

	errMsg := domain.ErrorMsg{
		ProcessId:    uuid.New().String(),
		Service:      service,
		SensorUuid:   sensorUuid,
		ErrorMsg:     errorMsg,
		ErrorDetails: errorDetails,
		PublishedOn:  now,
	}

	d.publishErrorMsg(errMsg, crosscutting.RoutingKeyLogsSeverityErrors.String())
}
