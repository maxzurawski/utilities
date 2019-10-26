package publishing

import (
	"time"

	"../crosscutting"
	"../domain"
	"github.com/google/uuid"
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
