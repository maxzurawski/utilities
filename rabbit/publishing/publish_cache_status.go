package publishing

import (
	"encoding/json"
	"time"

	"../crosscutting"
	"../domain"
	"github.com/labstack/gommon/log"
)

func (d *Publisher) PublishCacheUpdateStatus(
	processid,
	routingKey,
	service,
	cache,
	errorMsg string) {

	now := time.Now()

	msg := domain.CacheUpdateStatus{
		ProcessId:   processid,
		Service:     service,
		Cache:       cache,
		ErrorMsg:    errorMsg,
		PublishedOn: now,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Warn("could not marshal msg")
		return
	}

	d.Publish(
		crosscutting.TopicConfigurationChanged.String(),
		routingKey,
		string(bytes),
	)

	d.PublishExtInfo(msg)
}
