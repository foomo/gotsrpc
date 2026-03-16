package server

type ErrService string

const (
	ErrUnauthorized ErrService = "unauthorized"
)

type ScalarError string

func (s *ScalarError) String() string {
	return string(*s)
}

func (s *ScalarError) Error() string {
	return s.String()
}

type StructError struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (s *StructError) Error() string {
	return s.Message
}

type (
	ScalarA     string
	ScalarB     string
	MultiScalar struct {
		ScalarA `json:",omitempty,inline"`
		ScalarB `json:",omitempty,inline"`
	}
)

const (
	ScalarOne ScalarError = "one"
	ScalarTwo ScalarError = "two"

	ScalarAOne ScalarA = "one"
	ScalarATwo ScalarA = "two"

	ScalarBThree ScalarB = "three"
	ScalarBFour  ScalarB = "four"
)
