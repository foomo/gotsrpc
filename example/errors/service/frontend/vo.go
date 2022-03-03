package frontend

type ErrSimple string

type (
	ErrMulti struct {
		A ErrMultiA `json:"a,omitempty,union"`
		B ErrMultiB `json:"b,omitempty,union"`
	}
	ErrMultiA string
	ErrMultiB string
)
