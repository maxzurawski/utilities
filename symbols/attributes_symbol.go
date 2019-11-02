package symbols

type AttributeSymbol int

const (
	AcceptableMax AttributeSymbol = iota
	AcceptableMin
	Active
	NotificationAfterContinuousTransitionAmount
)

func (a AttributeSymbol) String() string {
	return [...]string{
		"ACCEPTABLE_MAX",
		"ACCEPTABLE_MIN",
		"ACTIVE",
		"NOTIFICATION_AFTER_CONTINUOUS_TRANSITION_AMOUNT",
	}[a]
}
