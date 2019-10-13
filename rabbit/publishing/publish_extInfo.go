package publishing

import (
	"github.com/xdevices/utilities/rabbit/crosscutting"
)

func (d *Publisher) PublishExtInfo(
	msg interface{}) {
	d.publishExtLogMsg(msg, crosscutting.RoutingKeyLogsSeverityInfo.String())
}
