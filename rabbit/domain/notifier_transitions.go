package domain

import "errors"

type NotifierTransitions int

const (
	FirstTransition NotifierTransitions = iota
	ContinuousTransition
	FinalTransition
)

var transitionNames = []string{"Initial transition", "Continuous transition", "Final transition"}

func (nt NotifierTransitions) String() string {
	return transitionNames[nt]
}

func NotifierTransitionFromName(name string) (NotifierTransitions, error) {
	switch name {
	case FirstTransition.String():
		return FirstTransition, nil
	case ContinuousTransition.String():
		return ContinuousTransition, nil
	case FinalTransition.String():
		return FinalTransition, nil
	default:
		return -1, errors.New("cannot recognize NotifierTransition")
	}
}
