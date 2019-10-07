package crosscutting

type RoutingKeys int

const (
	RoutingKeySensorTypes RoutingKeys = iota
	RoutingKeySensors
	RoutingKeyTemperatureMeasurement

	RoutingKeyLogsSeverityInfo
	RoutingKeyLogsSeverityErrors
	RoutingKeyLogsSeverityWarn

	RoutingKeyNotifierTemperatureMax
	RoutingKeyNotifierTemperatureMin
)

func (r RoutingKeys) String() string {
	return [...]string{
		"sensortypes", // RoutingKeySensorTypes
		"sensors",     // RoutingKeySensors
		"temperature", // RoutingKeyTemperatureMeasurement

		"info",  // RoutingKeyLogsSeverityInfo
		"error", // RoutingKeyLogsSeverityErrors
		"warn",  // RoutingKeyLogsSeverityWarn

		"temperature.max", // RoutingKeyNotifierTemperatureMax
		"temperature.min", // RoutingKeyNotifierTemperatureMin
	}[r]
}
