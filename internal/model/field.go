package model

type Field struct {
	Value    *Value
	Name     string    `json:",omitempty"`
	JSONInfo *JSONInfo `json:",omitempty"`
}
