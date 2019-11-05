package domain

type EmailDTO struct {
	Name       string   `json:"name"`
	Subject    string   `json:"subject"`
	Recipients []string `json:"recipients"`
	Msg        string   `json:"msg"`
}
