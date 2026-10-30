package gotsrpc

// ClientError reports a failure on the client side of a call: encoding the
// arguments, sending the request, a non-200 status or decoding the response.
// It wraps the underlying error.
type ClientError struct {
	error
}

// NewClientError returns a [ClientError] that wraps err.
func NewClientError(err error) *ClientError {
	return &ClientError{
		error: err,
	}
}

// Unwrap returns the wrapped error, or nil if e is nil.
func (e *ClientError) Unwrap() error {
	if e != nil && e.error != nil {
		return e.error
	}

	return nil
}
