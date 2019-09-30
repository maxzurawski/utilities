package resterror

type ErrorMsg struct {
	Msg string
}

func New(msg string) *ErrorMsg {
	return &ErrorMsg{Msg: msg}
}
