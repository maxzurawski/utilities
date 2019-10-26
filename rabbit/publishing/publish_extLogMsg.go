package publishing

import (
	"encoding/json"

	"../crosscutting"
	"github.com/labstack/gommon/log"
)

func (d *Publisher) publishExtLogMsg(
	logMsg interface{},
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
