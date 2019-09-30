package resterror

type ErrorMsg struct {
	Msg    string
	Detail string
}

func New(msg string) *ErrorMsg {
	return &ErrorMsg{Msg: msg, Detail: ""}
}
