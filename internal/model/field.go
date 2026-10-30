package model

// Field is a named, typed struct field, method argument or return value.
type Field struct {
	Value    *Value
	Name     string    `json:",omitempty"`
	JSONInfo *JSONInfo `json:",omitempty"`
}
