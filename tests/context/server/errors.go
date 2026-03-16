package server

type MyError struct {
	Payload string `json:"payload"`
}

func (e *MyError) Error() string {
	return e.Payload
}

func NewMyError(msg string) error {
	return &MyError{Payload: msg}
}
