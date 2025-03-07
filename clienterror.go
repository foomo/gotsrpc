package gotsrpc

type ClientError struct {
	error
}

func NewClientError(err error) *ClientError {
	return &ClientError{
		error: err,
	}
}

// Unwrap interface
func (e *ClientError) Unwrap() error {
	if e != nil && e.error != nil {
		return e.error
	}
	return nil
}
