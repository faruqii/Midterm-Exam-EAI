package services

type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}
