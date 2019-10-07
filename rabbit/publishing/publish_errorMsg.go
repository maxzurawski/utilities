package publishing

import (
	"encoding/json"

	"github.com/labstack/gommon/log"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/domain"
)

func (d *Publisher) publishErrorMsg(
	errMsg domain.ErrorMsg,
	routingKey string) {

	bytes, err := json.Marshal(errMsg)
	if err != nil {
		log.Warn("could not marshal msg")
		return
	}

	d.Publish(
		crosscutting.TopicLogs.String(),
		routingKey,
		string(bytes))
}
