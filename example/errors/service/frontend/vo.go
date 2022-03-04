package frontend

type ErrSimple string

type (
	ErrMulti struct {
		A ErrMultiA `json:"a,omitempty" gotsrpc:"union"`
		B ErrMultiB `json:"b,omitempty" gotsrpc:"union"`
	}
	ErrMultiA string
	ErrMultiB string
)
