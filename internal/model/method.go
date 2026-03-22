package model

type Method struct {
	Name   string
	Args   []*Field
	Return []*Field
}
