package model

// Method is a service method with its arguments and return values.
type Method struct {
	Name   string
	Args   []*Field
	Return []*Field
}
