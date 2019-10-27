package publishing

import (
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
)

func (d *Publisher) PublishExtInfo(
	msg interface{}) {
	d.publishExtLogMsg(msg, crosscutting.RoutingKeyLogsSeverityInfo.String())
}
