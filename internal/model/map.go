package model

// Map describes a Go map type.
type Map struct {
	// Key is the key type.
	Key *Value
	// Value is the element type.
	Value *Value
	// KeyType is the [ScalarType] of the key.
	KeyType string
	// KeyGoType is the Go name of the key type.
	KeyGoType string
}
