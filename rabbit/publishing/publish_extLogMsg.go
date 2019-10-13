package publishing

import (
	"encoding/json"

	"github.com/labstack/gommon/log"
	"github.com/xdevices/utilities/rabbit/crosscutting"
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
