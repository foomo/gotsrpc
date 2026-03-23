package model

type Array struct {
	Value *Value
	Len   int // 0 = slice, >0 = fixed-size array
}
