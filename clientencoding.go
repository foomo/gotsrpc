package gotsrpc

// ClientEncoding selects the wire format a [Client] uses for requests.
type ClientEncoding int

// Supported client encodings.
const (
	EncodingMsgpack = ClientEncoding(0)
	EncodingJson    = ClientEncoding(1)
)
