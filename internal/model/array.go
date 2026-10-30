package model

// Array describes a Go slice or array type.
type Array struct {
	// Value is the element type.
	Value *Value
	Len   int // 0 = slice, >0 = fixed-size array
}
