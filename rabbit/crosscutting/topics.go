package crosscutting

type Topics int

const (
	TopicConfigurationChanged Topics = iota
	TopicMeasurements
	TopicLogs
	TopicNotifier
)

func (c Topics) String() string {
	return [...]string{
		"xdevices.configuration.changed", // TopicConfigurationChanged
		"xdevices.measurements",          // TopicMeasurements
		"xdevices.logs",                  // TopicLogs
		"xdevices.notifier",              // TopicNotifier
	}[c]
}

func GetAllTopics() []string {
	return []string{
		TopicConfigurationChanged.String(),
		TopicMeasurements.String(),
		TopicLogs.String(),
		TopicNotifier.String()}
}
