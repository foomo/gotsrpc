package gotsrpc

type ClientError struct {
	error
}

func NewClientError(err error) *ClientError {
	return &ClientError{
		error: err,
	}
}
