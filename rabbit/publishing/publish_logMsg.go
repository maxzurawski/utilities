package publishing

import (
	"encoding/json"

	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/domain"
)

func (d *Publisher) publishLogMsg(
	logMsg domain.LogMsg,
	routingKey string) {

	bytes, err := json.Marshal(logMsg)
	if err != nil {
		log.Warn("could not marshal msg")
		return
	}

	d.Publish(
		crosscutting.TopicLogs.String(),
		routingKey,
		string(bytes))
}
