package model

// JSONInfo holds the options of a field's json and gotsrpc struct tags.
type JSONInfo struct {
	// Name is the JSON key, or empty to use the field name.
	Name string
	// Type is the TypeScript type override set by gotsrpc:"type:...".
	Type string
	// Union is set by gotsrpc:"union" and marks a union member.
	Union bool
	// Inline is set by the inline option.
	Inline bool
	// OmitEmpty is set by the omitempty option.
	OmitEmpty bool
	// Ignore is set by the "-" tag.
	Ignore bool
}
