package publishing

import (
	"../crosscutting"
)

func (d *Publisher) PublishExtInfo(
	msg interface{}) {
	d.publishExtLogMsg(msg, crosscutting.RoutingKeyLogsSeverityInfo.String())
}
