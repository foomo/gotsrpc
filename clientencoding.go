package gotsrpc

type ClientEncoding int

const (
	EncodingMsgpack = ClientEncoding(0)
	EncodingJson    = ClientEncoding(1) //nolint:staticcheck
)
