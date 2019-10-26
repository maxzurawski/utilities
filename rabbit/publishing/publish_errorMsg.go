package publishing

import (
	"encoding/json"

	"../crosscutting"
	"../domain"
	"github.com/labstack/gommon/log"
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
