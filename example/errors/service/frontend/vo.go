package frontend

type ErrSimple string

type (
	ErrMulti struct {
		A ErrMultiA `json:"a,union"`
		B ErrMultiB `json:"b,union"`
	}
	ErrMultiA string
	ErrMultiB string
)
