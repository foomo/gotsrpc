package gotsrpc

// errorReply is a marker wrapper that identifies error interface returns in response slices.
type errorReply struct{ err error }

// ErrorReply wraps an error return value so Reply can detect it at runtime.
// Used by generated proxy code for methods whose last return type is the error interface.
func ErrorReply(err error) any {
	return &errorReply{err: err}
}
