package symbols

type AttributeSymbol int

const (
	AcceptableMax AttributeSymbol = iota
	AcceptableMin
	Active
)

func (a AttributeSymbol) String() string {
	return [...]string{
		"ACCEPTABLE_MAX",
		"ACCEPTABLE_MIN",
		"ACTIVE",
	}[a]
}
